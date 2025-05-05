package postgres

import (
	"context"
	"gorm.io/gorm"
	"predictive-platform/internal/domain/model"
)

type leagueDB struct {
	db *gorm.DB
}

func NewLeagueDB(db *gorm.DB) *leagueDB {
	return &leagueDB{
		db: db,
	}
}

func (l *leagueDB) CreateLeague(ctx context.Context, league *model.League) (string, error) {
	err := l.db.WithContext(ctx).Create(league).Error
	if err != nil {
		return "", err
	}
	return league.ID, nil
}

func (l *leagueDB) GetLeagueByID(ctx context.Context, id string) (*model.League, error) {
	league := &model.League{}
	err := l.db.WithContext(ctx).Where("id = ?", id).First(league).Error
	if err != nil {
		return nil, err
	}
	return league, nil
}

func (l *leagueDB) GetAllLeagues(ctx context.Context) ([]model.League, error) {
	var league []model.League
	err := l.db.WithContext(ctx).Order("created_at DESC").Find(league).Error
	if err != nil {
		return nil, err
	}
	return league, nil
}
