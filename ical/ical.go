package ical

import (
	"fmt"
	"strconv"
	"time"

	"github.com/Olament/HLTV-Go/model"
	"github.com/gosimple/slug"
	"github.com/jordic/goics"
)

func getHltvMatchUrl(m *model.UpcomingMatch) string {
	return fmt.Sprintf(
		"https://hltv.org/matches/%s/%s-vs-%s-%s",
		strconv.Itoa(*m.ID),
		slug.Make(m.Team1.Name),
		slug.Make(m.Team2.Name),
		slug.Make(m.Event.Name),
	)
}

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
		s.AddProperty("DESCRIPTION", getHltvMatchUrl(match))
		c.AddComponent(s)
	}
	return c
}
