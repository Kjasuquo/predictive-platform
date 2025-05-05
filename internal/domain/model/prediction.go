package model

import (
	"predictive-platform/internal/domain/dto"
	"time"
)

type Prediction struct {
	Models
	UserID             string    `json:"user_id"`
	MatchID            string    `json:"match_id"`
	PredictedHomeScore int       `json:"predicted_home_score"`
	PredictedAwayScore int       `json:"predicted_away_score"`
	CreatedAt          time.Time `json:"created_at"`
	IsCorrect          bool      `json:"is_correct"`
	Points             int       `json:"points_awarded"`

	// Relationships
	User  User  `gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
	Match Match `gorm:"foreignKey:MatchID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE"`
}

func (p *Prediction) ToPredictionDTO() *dto.Prediction {
	out := &dto.Prediction{
		ID:                 p.ID,
		UserID:             p.UserID,
		MatchID:            p.MatchID,
		PredictedHomeScore: p.PredictedHomeScore,
		PredictedAwayScore: p.PredictedAwayScore,
		CreatedAt:          p.CreatedAt,
		IsCorrect:          p.IsCorrect,
		Points:             p.Points,
	}
	return out
}

func FromPredictiveDTO(p *dto.Prediction) *Prediction {
	model := Models{
		ID: p.ID,
	}

	response := &Prediction{
		Models:             model,
		UserID:             p.UserID,
		MatchID:            p.MatchID,
		PredictedHomeScore: p.PredictedHomeScore,
		PredictedAwayScore: p.PredictedAwayScore,
		CreatedAt:          p.CreatedAt,
		IsCorrect:          p.IsCorrect,
		Points:             p.Points,
	}

	return response
}
