package main

import (
	"bytes"
	"net/http"
	"strconv"

	hltv "github.com/Olament/HLTV-Go"
	"github.com/Olament/HLTV-Go/model"
	"github.com/gin-gonic/gin"
	"github.com/jordic/goics"

	"github.com/mammuth/hltv-api/ical"
)

func upcomingMatches(c *gin.Context) {
	matches, _ := UpcomingMatchesFromRequest(c)
	c.JSON(http.StatusOK, matches)
}

func upcomingMatchesIcal(c *gin.Context) {

	matches, _ := UpcomingMatchesFromRequest(c)

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
	c.HTML(http.StatusOK, "index.html", gin.H{})
}
func UpcomingMatchesFromRequest(c *gin.Context) (upcomingMatches []*model.UpcomingMatch, err error) {
	teamIDs := []int{}
	for _, idString := range c.QueryArray("team") {
		id, err := strconv.Atoi(idString)
		if err != nil {
			continue
		}
		teamIDs = append(teamIDs, id)
	}

	h := hltv.HLTV{
		Url:       "https://www.hltv.org",
		StaticURL: "",
	}
	return h.GetUpcomingMatches(hltv.UpcomingMatchesQuery{
		Team: teamIDs,
	})
}
