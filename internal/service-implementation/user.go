package service_implementation

import (
	"context"
	"errors"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"

	"predictive-platform/internal/domain/dto"
	"predictive-platform/internal/domain/model"
	"predictive-platform/internal/domain/repositories"
	"predictive-platform/pkg/config"
	"predictive-platform/pkg/jwt"
)

type UserClient struct {
	Config *config.Config
	UserDB repositories.UserDB
}

// NewUserClient constructor for User.
func NewUserClient(conf *config.Config, userDB repositories.UserDB) *UserClient {
	return &UserClient{Config: conf, UserDB: userDB}
}

func (u *UserClient) UserSignUp(ctx context.Context, user *dto.User) (userID, jwtToken string, err error) {

	userID, err = u.UserDB.CreateUser(ctx, model.FromUserDTO(user))
	if err != nil {
		return
	}
	jwtToken, err = jwt.GenerateToken(user.Email, time.Minute*20)
	if err != nil {
		log.Printf("critical error [generating token]: %v", err)
		return
	}

	return
}

func (u *UserClient) UserLogin(ctx context.Context, loginReq *dto.Login) (user *dto.User, jwtToken string, err error) {
	modelUser, err := u.UserDB.GetUserByEmail(ctx, loginReq.Email)
	if err != nil {
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(modelUser.PasswordHash), []byte(loginReq.Password))
	if err != nil {
		err = errors.New("wrong password")
		return
	}

	user = modelUser.ToUserDTO()

	jwtToken, err = jwt.GenerateToken(user.Email, jwt.AccessTokenValidity)
	if err != nil {
		return
	}

	return
}

func (u *UserClient) GetUserByEmail(ctx context.Context, email string) (user *dto.User, err error) {
	modelUser, err := u.UserDB.GetUserByEmail(ctx, email)
	if err != nil {
		return
	}

	user = modelUser.ToUserDTO()

	return
}

func (u *UserClient) GetUserByID(ctx context.Context, id string) (user *dto.User, err error) {
	modelUser, err := u.UserDB.GetUserByID(ctx, id)
	if err != nil {
		return
	}

	user = modelUser.ToUserDTO()

	return
}

func (u *UserClient) UpdatePassword(ctx context.Context, login *dto.Login) error {

	modelUser, err := u.UserDB.GetUserByEmail(ctx, login.Email)
	if err != nil {
		return err
	}

	dtoUser := modelUser.ToUserDTO()
	dtoUser.Password = login.Password

	err = u.UserDB.UpdateUser(ctx, model.FromUserDTO(dtoUser))
	if err != nil {
		return err
	}
	return nil
}
