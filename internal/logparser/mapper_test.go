package logparser

import (
	"encoding/json"
	"testing"
)

func TestMatchesToGameRankingJSON(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		matches  []*Match
		expected string
		wantErr  bool
	}{
		{
			name: "Single match with players and kills",
			matches: []*Match{
				{
					ID:         1,
					TotalKills: 5,
					Players:    map[string]struct{}{"Player1": {}, "Player2": {}},
					Kills:      map[string]int{"Player1": 3, "Player2": 2},
				},
			},
			expected: `{
	"game_1": {
		"total_kills": 5,
		"players": [
			"Player1",
			"Player2"
		],
		"kills": {
			"Player1": 3,
			"Player2": 2
		}
	}
}`,
			wantErr: false,
		},
		{
			name: "Multiple matches with players and kills",
			matches: []*Match{
				{
					ID:         1,
					TotalKills: 5,
					Players:    map[string]struct{}{"Player1": {}, "Player2": {}},
					Kills:      map[string]int{"Player1": 3, "Player2": 2},
				},
				{
					ID:         2,
					TotalKills: 8,
					Players:    map[string]struct{}{"Player3": {}, "Player4": {}},
					Kills:      map[string]int{"Player3": 5, "Player4": 3},
				},
			},
			expected: `{
	"game_1": {
		"total_kills": 5,
		"players": [
			"Player1",
			"Player2"
		],
		"kills": {
			"Player1": 3,
			"Player2": 2
		}
	},
	"game_2": {
		"total_kills": 8,
		"players": [
			"Player3",
			"Player4"
		],
		"kills": {
			"Player3": 5,
			"Player4": 3
		}
	}
}`,
			wantErr: false,
		},
		{
			name:     "Empty matches",
			matches:  []*Match{},
			expected: "{}",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MatchesToGameRankingJSON(tt.matches)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatchesToGameRankingJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var gotJSON map[string]interface{}
			var expectedJSON map[string]interface{}

			err = json.Unmarshal([]byte(got), &gotJSON)
			if err != nil {
				t.Fatalf("Failed to unmarshal got JSON: %v", err)
			}

			err = json.Unmarshal([]byte(tt.expected), &expectedJSON)
			if err != nil {
				t.Fatalf("Failed to unmarshal expected JSON: %v", err)
			}

			if !compareJSON(gotJSON, expectedJSON) {
				t.Errorf("MatchesToGameRankingJSON() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func TestMatchesToGameDeathCausesJSON(t *testing.T) {
	// Define test cases
	tests := []struct {
		name     string
		matches  []*Match
		expected string
		wantErr  bool
	}{
		{
			name: "Single match with kill causes",
			matches: []*Match{
				{
					ID:           1,
					KillsByMeans: map[string]int{"MOD_ROCKET": 3, "MOD_RAILGUN": 2},
				},
			},
			expected: `{
	"game_1": {
		"kills_by_means": {
			"MOD_ROCKET": 3,
			"MOD_RAILGUN": 2
		}
	}
}`,
			wantErr: false,
		},
		{
			name: "Multiple matches with kill causes",
			matches: []*Match{
				{
					ID:           1,
					KillsByMeans: map[string]int{"MOD_ROCKET": 3, "MOD_RAILGUN": 2},
				},
				{
					ID:           2,
					KillsByMeans: map[string]int{"MOD_SHOTGUN": 5, "MOD_PLASMA": 3},
				},
			},
			expected: `{
	"game_1": {
		"kills_by_means": {
			"MOD_ROCKET": 3,
			"MOD_RAILGUN": 2
		}
	},
	"game_2": {
		"kills_by_means": {
			"MOD_SHOTGUN": 5,
			"MOD_PLASMA": 3
		}
	}
}`,
			wantErr: false,
		},
		{
			name:     "Empty matches",
			matches:  []*Match{},
			expected: "{}",
			wantErr:  false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := MatchesToGameDeathCausesJSON(tt.matches)
			if (err != nil) != tt.wantErr {
				t.Errorf("MatchesToGameDeathCausesJSON() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			var gotJSON map[string]interface{}
			var expectedJSON map[string]interface{}

			err = json.Unmarshal([]byte(got), &gotJSON)
			if err != nil {
				t.Fatalf("Failed to unmarshal got JSON: %v", err)
			}

			err = json.Unmarshal([]byte(tt.expected), &expectedJSON)
			if err != nil {
				t.Fatalf("Failed to unmarshal expected JSON: %v", err)
			}

			if !compareJSON(gotJSON, expectedJSON) {
				t.Errorf("MatchesToGameDeathCausesJSON() = %v, expected %v", got, tt.expected)
			}
		})
	}
}

func compareJSON(got, expected map[string]interface{}) bool {
	gotBytes, err := json.Marshal(got)
	if err != nil {
		return false
	}
	expectedBytes, err := json.Marshal(expected)
	if err != nil {
		return false
	}
	return string(gotBytes) == string(expectedBytes)
}
