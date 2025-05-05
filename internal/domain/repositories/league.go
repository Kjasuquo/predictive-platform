package repositories

import (
	"context"
	"predictive-platform/internal/domain/model"
)

type LeagueDB interface {
	CreateLeague(ctx context.Context, league *model.League) (string, error)
	GetLeagueByID(ctx context.Context, id string) (*model.League, error)
	GetAllLeagues(ctx context.Context) ([]model.League, error)
}
