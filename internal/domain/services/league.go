package services

import (
	"context"
	"predictive-platform/internal/domain/dto"
)

type LeagueService interface {
	CreateLeague(ctx context.Context, league *dto.League) (string, error)
	GetLeagueByID(ctx context.Context, id string) (*dto.League, error)
	GetAllLeagues(ctx context.Context) ([]dto.League, error)
}
