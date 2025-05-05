package dto

import "time"

type League struct {
	ID        string    `json:"id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	Name      string    `json:"name"`
}

type LeagueMember struct {
	ID       string    `json:"id"`
	Joined   time.Time `json:"joined,omitempty"`
	LeagueID string    `json:"league_id"`
	UserID   string    `json:"user_id" gorm:"unique"`
}
