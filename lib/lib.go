package lib

import (
	"fmt"
	"strconv"

	hltv "github.com/Olament/HLTV-Go"
	"github.com/Olament/HLTV-Go/model"
	"github.com/gin-gonic/gin"
	"github.com/gosimple/slug"
)

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

func GetHltvMatchUrl(m *model.UpcomingMatch) string {
	return fmt.Sprintf(
		"https://hltv.org/matches/%s/%s-vs-%s-%s",
		strconv.Itoa(*m.ID),
		slug.Make(m.Team1.Name),
		slug.Make(m.Team2.Name),
		slug.Make(m.Event.Name),
	)
}
