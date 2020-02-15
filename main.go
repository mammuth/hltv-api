package main

import (
	"github.com/gin-gonic/gin"
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
