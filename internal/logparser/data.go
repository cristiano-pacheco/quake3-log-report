package logparser

type LogEvent struct {
	Timestamp string
	EventType string
	Killer    string
	Killed    string
	KillType  string
}

type Match struct {
	ID           int
	TotalKills   int
	Players      map[string]struct{}
	Kills        map[string]int
	KillsByMeans map[string]int
}
