package logparser

import (
	"testing"
)

func TestParseQuakeLog(t *testing.T) {
	// Define the path to the test log file
	testLogPath := "../../tests/test.log"

	// Call the function to test
	matches, err := ParseQuakeLog(testLogPath)
	if err != nil {
		t.Fatalf("Unexpected error: %v", err)
	}

	// Define expected results based on the log content
	expectedMatches := []*Match{
		{
			ID:           1,
			TotalKills:   5,
			Players:      map[string]struct{}{"Isgalamido": {}, "Mocinha": {}},
			Kills:        map[string]int{"Isgalamido": 4, "Mocinha": 1},
			KillsByMeans: map[string]int{"MOD_ROCKET_SPLASH": 5},
		},
		{
			ID:           2,
			TotalKills:   12,
			Players:      map[string]struct{}{"Isgalamido": {}, "Mocinha": {}},
			Kills:        map[string]int{"Isgalamido": 4, "Mocinha": 8},
			KillsByMeans: map[string]int{"MOD_ROCKET_SPLASH": 12},
		},
		{
			ID:           3,
			TotalKills:   4,
			Players:      map[string]struct{}{"Isgalamido": {}, "Mocinha": {}},
			Kills:        map[string]int{"Isgalamido": 0, "Mocinha": 0},
			KillsByMeans: map[string]int{"MOD_ROCKET_SPLASH": 2, "MOD_TRIGGER_HURT": 2},
		},
		{
			ID:           4,
			TotalKills:   2,
			Players:      map[string]struct{}{"Isgalamido": {}, "Mocinha": {}},
			Kills:        map[string]int{"Isgalamido": -1, "Mocinha": -1},
			KillsByMeans: map[string]int{"MOD_TRIGGER_HURT": 2},
		},
	}

	// Check if the length of the matches slices are equal
	if len(matches) != len(expectedMatches) {
		t.Fatalf("Expected %d matches, but got %d", len(expectedMatches), len(matches))
	}

	// Compare each match in the slices
	for i, match := range matches {
		matchNumber := i + 1
		expected := expectedMatches[i]

		if match.ID != expected.ID {
			t.Errorf("Match %d: Expected match ID %d, but got %d", matchNumber, expected.ID, match.ID)
		}

		if match.TotalKills != expected.TotalKills {
			t.Errorf("Match %d: Expected TotalKills %d, but got %d", matchNumber, expected.TotalKills, match.TotalKills)
		}

		// Compare Players map
		if len(match.Players) != len(expected.Players) {
			t.Errorf("Match %d: Expected %d players, but got %d", matchNumber, len(expected.Players), len(match.Players))
		} else {
			for player := range expected.Players {
				if _, exists := match.Players[player]; !exists {
					t.Errorf("Match %d: Expected player %s in match, but not found", matchNumber, player)
				}
			}
		}

		// Compare Kills map
		if len(match.Kills) != len(expected.Kills) {
			t.Errorf("Match %d: Expected %d kills, but got %d", matchNumber, len(expected.Kills), len(match.Kills))
		} else {
			for player, kills := range expected.Kills {
				if match.Kills[player] != kills {
					t.Errorf("Match %d: Expected %d kills for player %s, but got %d", matchNumber, kills, player, match.Kills[player])
				}
			}
		}

		// Compare KillsByMeans map
		if len(match.KillsByMeans) != len(expected.KillsByMeans) {
			t.Errorf("Match %d: Expected %d kills by means, but got %d", matchNumber, len(expected.KillsByMeans), len(match.KillsByMeans))
		} else {
			for means, kills := range expected.KillsByMeans {
				if match.KillsByMeans[means] != kills {
					t.Errorf("Match %d: Expected %d kills by %s, but got %d", matchNumber, kills, means, match.KillsByMeans[means])
				}
			}
		}
	}
}
