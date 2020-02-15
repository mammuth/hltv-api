package main

import (
	"bytes"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jordic/goics"

	"github.com/mammuth/hltv-api/ical"
	"github.com/mammuth/hltv-api/lib"
)

func main() {
	gin.ForceConsoleColor()

	router := gin.Default()
	router.Use(gin.Recovery())

	router.LoadHTMLGlob("templates/*")
	router.Static("/assets", "./assets")

	router.GET("/", index)
	api := router.Group("/api")
	{
		api.GET("/upcoming-matches", upcomingMatches)
		api.GET("/upcoming-matches/ical", upcomingMatchesIcal)
	}

	router.Run()
}

func upcomingMatches(c *gin.Context) {
	matches, _ := lib.UpcomingMatchesFromRequest(c)
	c.JSON(http.StatusOK, matches)
}

func upcomingMatchesIcal(c *gin.Context) {

	matches, _ := lib.UpcomingMatchesFromRequest(c)

	var icalBytes bytes.Buffer
	generatedIcal := ical.UpcomingMatchesICal(matches)
	icalWriter := goics.NewICalEncode(&icalBytes)
	generatedIcal.Write(icalWriter)

	header := c.Writer.Header()
	header.Set("Content-type", "text/calendar")
	header.Set("charset", "utf-8")
	header.Set("Content-Disposition", "inline")
	header.Set("filename", "calendar.ics")
	c.Data(http.StatusOK, "text/calendar", icalBytes.Bytes())
}

func index(c *gin.Context) {
	// c.String(200, "OKOK")
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
