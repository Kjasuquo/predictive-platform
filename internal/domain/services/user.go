package services

import (
	"context"

	"predictive-platform/internal/domain/dto"
)

type UserService interface {
	UserSignUp(ctx context.Context, user *dto.User) (userID, jwtToken string, err error)
	UserLogin(ctx context.Context, loginReq *dto.Login) (user *dto.User, jwtToken string, err error)
	GetUserByEmail(ctx context.Context, email string) (user *dto.User, err error)
	GetUserByID(ctx context.Context, id string) (user *dto.User, err error)
	UpdatePassword(ctx context.Context, login *dto.Login) error
}
