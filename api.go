package main // import "github.com/dayvillefire/iardisplay"

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/dayvillefire/iardisplay/config"
	"github.com/gin-gonic/gin"
	"github.com/jbuchbinder/cadmonitor/monitor"
	"github.com/jbuchbinder/iarapi"
)

func initAPI(m *gin.Engine) {
	g := m.Group("/api")

	g.GET("/config", apiUIConfig)

	c := g.Group("/cad")
	c.GET("/current", apiCadCurrent)
	c.GET("/cleared/:dt", apiCadClearedDate)

	i := g.Group("/iar")
	i.GET("/incidents", apiIarIncidents)
	i.GET("/incident/:id", apiIarIncidentDetail)
	i.GET("/messages", apiIarMessages)
	i.GET("/responding", apiIarNowResponding)
	i.GET("/schedule", apiIarSchedule)
}

func apiCadCurrent(c *gin.Context) {
	a, err := cad.GetActiveCalls()
	if config.Config.Debug {
		log.Printf("apiCadCurrent : CAD CURRENT : %#v", a)
	}
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	cs := map[string]monitor.CallStatus{}
	for _, x := range a {
		detail, err := cadStatusCache.RetrieveWithCache(x)
		if err == nil {
			cs[detail.DispatchTime.Format(OurTimeFormat)] = detail
		} else {
			log.Printf("apiCadCurrent: ERROR: %#v", err)
		}
	}
	c.JSON(http.StatusOK, cs)
}

func apiCadClearedDate(c *gin.Context) {
	a, err := cad.GetClearedCalls(strings.ReplaceAll(c.Param("dt"), "-", "/"))
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	cs := map[string]monitor.CallStatus{}
	for _, x := range a {
		detail, err := cadStatusCache.RetrieveWithCache(x)
		if err == nil {
			cs[detail.DispatchTime.Format(OurTimeFormat)] = detail
		}
	}
	c.JSON(http.StatusOK, cs)
}

func apiIarIncidents(c *gin.Context) {
	//i, err := iar.GetLatestIncidents()
	iraw, err := iarCache.RetrieveWithCache(IarLatestIncidents, "")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	i := iraw.([]iarapi.IncidentInfoData)
	d := map[string]iarapi.IncidentInfoData{}
	for _, x := range i {
		detail, err := iar.GetIncidentInfo(x.ID)
		if err == nil {
			d[fmt.Sprintf("%d", x.ID)] = detail
		}
	}
	c.JSON(http.StatusOK, gin.H{"incidents": i, "detail": d})
}

func apiIarIncidentDetail(c *gin.Context) {
	//id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	//i, err := iar.GetIncidentInfo(id)
	i, err := iarCache.RetrieveWithCache(IarIncidentInfoData, c.Param("id"))
	if err != nil { ////a
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, i.([]iarapi.IncidentInfoData))
}

func apiIarMessages(c *gin.Context) {
	ms, err := iarCache.RetrieveWithCache(IarDispatchMessage, "")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, ms.([]iarapi.DispatchMessage))
}

func apiIarNowResponding(c *gin.Context) {
	nr, err := iar.GetNowRespondingWithSort()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nr)
}

func apiIarSchedule(c *gin.Context) {
	s, err := iarCache.RetrieveWithCache(IarOnSchedule, "")
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, s.([]iarapi.OnSchedule))
}

func apiUIConfig(c *gin.Context) {
	c.JSON(http.StatusOK, map[string]interface{}{
		"debug":             config.Config.Debug,
		"unitSuffix":        config.Config.Accounts.Cad.UnitSuffix,
		"iarAcceptPatterns": config.Config.Accounts.Iar.AcceptPatterns,
		"ignorePatterns":    config.Config.Accounts.Cad.IgnorePatterns,
	})
}

func contextRequestHeader(c *gin.Context, key string) string {
	if config.Config.Debug {
		log.Printf("contextRequestHeader: %#v", c.Request.Header[key])
	}
	if values, _ := c.Request.Header[key]; len(values) > 0 {
		return values[0]
	}
	return ""
}
