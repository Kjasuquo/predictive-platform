package repositories

import (
	"context"

	"predictive-platform/internal/domain/model"
)

type UserDB interface {
	CreateUser(ctx context.Context, user *model.User) (string, error)
	GetUserByEmail(ctx context.Context, email string) (*model.User, error)
	GetUserByID(ctx context.Context, id string) (*model.User, error)
	UpdateUser(ctx context.Context, user *model.User) error
}
