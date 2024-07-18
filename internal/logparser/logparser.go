// Package logparser provides functions to parse Quake 3 Arena log files.
package logparser

import (
	"bufio"
	"os"
	"regexp"
	"strings"
)

func ParseQuakeLog(filepath string) ([]*Match, error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Initialize a scanner to read the file line by line
	scanner := bufio.NewScanner(file)
	var matches []*Match
	var currentMatch *Match

	// Process each line of the file
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		if isInitGame(line) {
			// New match starts
			currentMatch = &Match{
				ID:           len(matches) + 1,
				Players:      make(map[string]struct{}),
				Kills:        make(map[string]int),
				KillsByMeans: make(map[string]int),
			}
			matches = append(matches, currentMatch)
		} else if isShutdownGame(line) {
			// Match ends
			currentMatch = nil
		} else if currentMatch != nil {
			// If inside a match, check for kill events
			event := parseKillEvent(line)
			if event != nil {
				currentMatch.TotalKills++
				if event.Killer != "<world>" {
					currentMatch.Players[event.Killer] = struct{}{}
					currentMatch.Kills[event.Killer]++
				}
				currentMatch.Players[event.Killed] = struct{}{}
				if event.Killer == "<world>" {
					currentMatch.Kills[event.Killed]--
				}
				currentMatch.KillsByMeans[event.KillType]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return matches, nil
}

// Function to detect the start of a match
func isInitGame(line string) bool {
	return regexp.MustCompile(`^(\d+:\d+)\s+InitGame:`).MatchString(line)
}

// Function to detect the end of a match
func isShutdownGame(line string) bool {
	return regexp.MustCompile(`^(\d+:\d+)\s+ShutdownGame:`).MatchString(line)
}

// Function to parse lines related to kill events
func parseKillEvent(line string) *LogEvent {
	// Adjusted regex pattern to capture the kill event line correctly
	re := regexp.MustCompile(`^(\d+:\d+)\s+Kill:\s+\d+\s+\d+\s+\d+:\s+(.+?)\s+killed\s+(.+?)\s+by\s+(\w+)$`)
	matches := re.FindStringSubmatch(line)
	if matches == nil {
		return nil
	}

	return &LogEvent{
		Timestamp: matches[1],
		EventType: "Kill",
		Killer:    matches[2],
		Killed:    matches[3],
		KillType:  matches[4],
	}
}
