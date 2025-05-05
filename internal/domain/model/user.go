package model

import (
	"golang.org/x/crypto/bcrypt"
	"log"

	"predictive-platform/internal/domain/dto"
)

type User struct {
	Models
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique"`
	PasswordHash string `json:"password_hash"`
	Points       int    `json:"points"`
}

func (u *User) HashPassword(password string) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.PasswordHash = string(hashedPassword)
	return nil
}

func (u *User) ToUserDTO() *dto.User {
	out := &dto.User{
		ID:     u.ID,
		Name:   u.Name,
		Email:  u.Email,
		Points: u.Points,
	}
	return out
}

func FromUserDTO(u *dto.User) *User {
	model := Models{
		ID: u.ID,
	}

	response := &User{
		Models: model,
		Name:   u.Name,
		Email:  u.Email,
		Points: u.Points,
	}

	if u.Password != "" {
		err := response.HashPassword(u.Password)
		if err != nil {
			log.Println("cannot hash password")
			return nil
		}
	}

	return response
}
