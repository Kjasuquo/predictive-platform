package dto

import (
	"time"
)

type Match struct {
	ID         string    `json:"id"`
	HomeTeam   string    `json:"home_team"`
	AwayTeam   string    `json:"away_team"`
	StartTime  time.Time `json:"start_time"`
	Status     string    `json:"status"`
	HomeScore  int       `json:"home_score"`
	AwayScore  int       `json:"away_score"`
	LeagueID   string    `json:"league_id"`
	LeagueName string    `json:"league_name"`
}
