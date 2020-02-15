package ical

import (
	"fmt"
	"time"

	"github.com/Olament/HLTV-Go/model"
	"github.com/jordic/goics"

	"github.com/mammuth/hltv-api/lib"
)

func UpcomingMatchesICal(matches []*model.UpcomingMatch) goics.Componenter {
	c := goics.NewComponent()
	c.SetType("VCALENDAR")
	c.AddProperty("CALSCAL", "GREGORIAN")

	for _, match := range matches {
		s := goics.NewComponent()
		s.SetType("VEVENT")

		k, v := goics.FormatDateTimeField("DTSTART", match.Date)
		s.AddProperty(k, v)
		k, v = goics.FormatDateTimeField("DTEND", match.Date.Add(time.Hour)) // We set the match duration to 1 hour
		s.AddProperty(k, v)

		s.AddProperty("SUMMARY", fmt.Sprintf("%s vs %s", match.Team1.Name, match.Team2.Name))
		s.AddProperty("DESCRIPTION", lib.GetHltvMatchUrl(match))
		c.AddComponent(s)
	}
	return c
}
