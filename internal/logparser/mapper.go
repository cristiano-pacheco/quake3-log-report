package logparser

import (
	"encoding/json"
	"fmt"
)

func MatchesToGameRankingJSON(matches []*Match) (string, error) {
	type matchData struct {
		TotalKills int            `json:"total_kills"`
		Players    []string       `json:"players"`
		Kills      map[string]int `json:"kills"`
	}

	ranking := make(map[string]*matchData)

	for _, match := range matches {
		m := matchData{
			TotalKills: match.TotalKills,
			Players:    make([]string, 0, len(match.Players)),
			Kills:      make(map[string]int),
		}

		for player := range match.Players {
			m.Players = append(m.Players, player)
		}

		for player, kills := range match.Kills {
			m.Kills[player] = kills
		}

		id := fmt.Sprintf("game_%d", match.ID)
		ranking[id] = &m
	}

	jsonData, err := json.MarshalIndent(ranking, "", "\t")
	if err != nil {
		return "", nil
	}

	return string(jsonData), nil
}

func MatchesToGameDeathCausesJSON(matches []*Match) (string, error) {
	type data struct {
		KillByMeans map[string]int `json:"kills_by_means"`
	}

	deathCauses := make(map[string]*data)

	for _, match := range matches {
		d := data{
			KillByMeans: make(map[string]int),
		}

		for cause, kills := range match.KillsByMeans {
			d.KillByMeans[cause] = kills
		}

		id := fmt.Sprintf("game_%d", match.ID)
		deathCauses[id] = &d
	}

	jsonData, err := json.MarshalIndent(deathCauses, "", "\t")
	if err != nil {
		return "", nil
	}

	return string(jsonData), nil
}
