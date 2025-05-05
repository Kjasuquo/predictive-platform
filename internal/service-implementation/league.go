package service_implementation

import (
	"context"
	"predictive-platform/internal/domain/dto"
	"predictive-platform/internal/domain/model"
	"predictive-platform/internal/domain/repositories"
	"predictive-platform/pkg/config"
)

type LeagueClient struct {
	Config   *config.Config
	LeagueDB repositories.LeagueDB
}

func NewLeagueClient(conf *config.Config, leagueDB repositories.LeagueDB) *LeagueClient {
	return &LeagueClient{Config: conf, LeagueDB: leagueDB}
}

func (l *LeagueClient) CreateLeague(ctx context.Context, league *dto.League) (string, error) {

	id, err := l.LeagueDB.CreateLeague(ctx, model.FromLeagueDTO(league))
	if err != nil {
		return "", err
	}

	return id, nil
}

func (l *LeagueClient) GetLeagueByID(ctx context.Context, id string) (*dto.League, error) {
	league, err := l.LeagueDB.GetLeagueByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return league.ToLeagueDTO(), nil
}

func (l *LeagueClient) GetAllLeagues(ctx context.Context) ([]dto.League, error) {
	league, err := l.LeagueDB.GetAllLeagues(ctx)
	if err != nil {
		return nil, err
	}
	var dtoLeagues []dto.League
	for _, ml := range league {
		leag := ml.ToLeagueDTO()
		dtoLeagues = append(dtoLeagues, *leag)
	}

	return dtoLeagues, nil
}
