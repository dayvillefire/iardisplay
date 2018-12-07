package main // import "github.com/dayvillefire/iardisplay"

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/dayvillefire/iardisplay/config"
	"github.com/gin-gonic/gin"
	"github.com/jbuchbinder/iarapi"
)

type stringarray []string

type genericReturnMap map[string]interface{}

func initApi(m *gin.Engine) {
	g := m.Group("/api")

	g.GET("/incidents", apiIncidents)
	g.GET("/incident/:id", apiIncidentDetail)
	g.GET("/messages", apiMessages)
	g.GET("/responding", apiNowResponding)
	g.GET("/schedule", apiSchedule)
}

func apiIncidents(c *gin.Context) {
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

func apiIncidentDetail(c *gin.Context) {
	id, _ := strconv.ParseInt(c.Param("id"), 10, 64)
	i, err := iar.GetIncidentInfo(id)
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, i)
}

func apiMessages(c *gin.Context) {
	ms, err := iar.ListWithParser()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, ms)
}

func apiNowResponding(c *gin.Context) {
	nr, err := iar.GetNowRespondingWithSort()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, nr)
}

func apiSchedule(c *gin.Context) {
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
