package model

import (
	"predictive-platform/internal/domain/dto"
	"time"
)

const (
	StatusUpcoming = "upcoming"
	StatusLive     = "live"
	StatusFinished = "finished"
)

type Match struct {
	Models
	HomeTeam  string    `json:"home_team"`
	AwayTeam  string    `json:"away_team"`
	StartTime time.Time `json:"start_time"`
	Status    string    `json:"status" gorm:"type:varchar(20);default:'upcoming'"`
	HomeScore int       `json:"home_score"`
	AwayScore int       `json:"away_score"`

	LeagueID string `json:"league_id"`
	League   League `gorm:"foreignKey:LeagueID;constraint:OnUpdate:CASCADE,OnDelete:SET NULL"`
}

func (m *Match) ToMatchDTO() *dto.Match {
	out := &dto.Match{
		ID:         m.ID,
		HomeTeam:   m.HomeTeam,
		AwayTeam:   m.AwayTeam,
		StartTime:  m.StartTime,
		Status:     m.Status,
		HomeScore:  m.HomeScore,
		AwayScore:  m.AwayScore,
		LeagueID:   m.LeagueID,
		LeagueName: m.League.Name,
	}
	return out
}

func FromMatchDTO(m *dto.Match) *Match {
	model := Models{
		ID: m.ID,
	}

	response := &Match{
		Models:    model,
		HomeTeam:  m.HomeTeam,
		AwayTeam:  m.AwayTeam,
		StartTime: m.StartTime,
		Status:    m.Status,
		HomeScore: m.HomeScore,
		AwayScore: m.AwayScore,
		LeagueID:  m.LeagueID,
	}

	return response
}
