package model

import (
	"predictive-platform/internal/domain/dto"
)

type League struct {
	Models
	Name string `json:"name"`
}

type LeagueMember struct {
	Models
	LeagueID string `json:"league_id"`
	UserID   string `json:"user_id"`

	League League `gorm:"foreignKey:LeagueID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	User   User   `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (l *League) ToLeagueDTO() *dto.League {
	out := &dto.League{
		ID:        l.ID,
		Name:      l.Name,
		CreatedAt: l.CreatedAt,
	}
	return out
}

func FromLeagueDTO(l *dto.League) *League {
	model := Models{
		ID:        l.ID,
		CreatedAt: l.CreatedAt,
	}

	response := &League{
		Models: model,
		Name:   l.Name,
	}

	return response
}

func (l *LeagueMember) ToLeagueMembersDTO() *dto.LeagueMember {
	out := &dto.LeagueMember{
		ID:       l.ID,
		LeagueID: l.LeagueID,
		UserID:   l.UserID,
		Joined:   l.CreatedAt,
	}
	return out
}

func FromLeagueMembersDTO(l *dto.LeagueMember) *LeagueMember {
	model := Models{
		ID:        l.ID,
		CreatedAt: l.Joined,
	}

	response := &LeagueMember{
		Models:   model,
		LeagueID: l.LeagueID,
		UserID:   l.UserID,
	}

	return response
}
