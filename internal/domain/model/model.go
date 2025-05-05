package model

import (
	"gorm.io/gorm"
	"time"

	"github.com/google/uuid"
)

type Models struct {
	ID        string    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"created_at,omitempty"`
	UpdatedAt time.Time `gorm:"autoUpdateTime" json:"updated_at,omitempty"`
}

func (u *Models) BeforeCreate(tx *gorm.DB) (err error) {
	if u.ID == "" {
		u.ID = uuid.New().String()
	}
	u.CreatedAt = time.Now()
	u.UpdatedAt = u.CreatedAt
	return nil
}
