package main // import "github.com/dayvillefire/iardisplay"

import (
	"log"
	"net/http"

	"github.com/dayvillefire/iardisplay/config"
	"github.com/gin-gonic/gin"
)

type stringarray []string

type genericReturnMap map[string]interface{}

func initApi(m *gin.Engine) {
	g := m.Group("/api")

	g.GET("/messages", apiMessages)
	g.GET("/responding", apiNowResponding)
	g.GET("/schedule", apiSchedule)
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
