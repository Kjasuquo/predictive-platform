package dto

import "time"

type Prediction struct {
	ID                 string    `json:"id"`
	UserID             string    `json:"user_id"`
	MatchID            string    `json:"match_id"`
	PredictedHomeScore int       `json:"predicted_home_score"`
	PredictedAwayScore int       `json:"predicted_away_score"`
	CreatedAt          time.Time `json:"created_at"`
	IsCorrect          bool      `json:"is_correct"`
	Points             int       `json:"points_awarded"`
}
