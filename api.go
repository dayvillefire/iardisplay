package main // import "github.com/dayvillefire/iardisplay"

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/dayvillefire/iardisplay/config"
	"github.com/gin-gonic/gin"
	"github.com/jbuchbinder/cadmonitor/monitor"
	"github.com/jbuchbinder/iarapi"
)

func initApi(m *gin.Engine) {
	g := m.Group("/api")

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
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	cs := map[string]monitor.CallStatus{}
	for _, x := range a {
		detail, err := cad.GetStatus(x)
		if err == nil {
			cs[detail.DispatchTime.Format("2006-01-02 15:04:05")] = detail
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
		detail, err := cad.GetStatus(x)
		if err == nil {
			cs[detail.DispatchTime.Format("2006-01-02 15:04:05")] = detail
		}
	}
	c.JSON(http.StatusOK, cs)
}

func apiIarIncidents(c *gin.Context) {
	i, err := iar.GetLatestIncidents()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	d := map[string]iarapi.IncidentInfoData{}
	for _, x := range i {
		detail, err := iar.GetIncidentInfo(x.Id)
		if err == nil {
			d[fmt.Sprintf("%d", x.Id)] = detail
		}
	}
	c.JSON(http.StatusOK, gin.H{"incidents": i, "detail": d})
}

func apiIarIncidentDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	i, err := iar.GetIncidentInfo(id)
	if err != nil { ////a
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, i)
}

func apiIarMessages(c *gin.Context) {
	ms, err := iar.ListWithParser()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, ms)
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
	s, err := iar.GetOnScheduleWithSort()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, s)
}

func ContextRequestHeader(c *gin.Context, key string) string {
	if config.Config.Debug {
		log.Printf("ContextRequestHeader: %#v", c.Request.Header[key])
	}
	if values, _ := c.Request.Header[key]; len(values) > 0 {
		return values[0]
	}
	return ""
}
