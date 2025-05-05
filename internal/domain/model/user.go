package model

import (
	"golang.org/x/crypto/bcrypt"
	"log"

	"predictive-platform/internal/domain/dto"
)

const (
	Admin       = "admin"
	RegularUser = "user"
)

// User represents an app user
type User struct {
	Models
	Name         string `json:"name"`
	Email        string `json:"email" gorm:"unique"`
	PasswordHash string `json:"password_hash"`
	Image        string `json:"image"`
	Admin        bool   `json:"admin"`
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
		ID:    u.ID,
		Name:  u.Name,
		Email: u.Email,
		Admin: u.Admin,
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
		Admin:  u.Admin,
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
