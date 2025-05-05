package postgres

import (
	"context"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"predictive-platform/internal/domain/model"
)

type userDB struct {
	db *gorm.DB
}

func NewUserDB(db *gorm.DB) *userDB {
	return &userDB{
		db: db,
	}
}

func (u *userDB) CreateUser(ctx context.Context, user *model.User) (string, error) {
	err := u.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return "", err
	}
	return user.ID, nil
}

func (u *userDB) GetUserByEmail(ctx context.Context, email string) (*model.User, error) {
	user := &model.User{}
	err := u.db.WithContext(ctx).Where("email = ?", email).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userDB) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	user := &model.User{}
	err := u.db.WithContext(ctx).Where("id = ?", id).First(user).Error
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (u *userDB) UpdateUser(ctx context.Context, user *model.User) error {
	// Perform the update
	result := u.db.WithContext(ctx).
		Model(&model.User{}).
		Where("id = ?", user.ID).
		Updates(user)

	// Check for errors
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return fmt.Errorf("email already exists: %w", result.Error)
		}
		return fmt.Errorf("failed to update user: %w", result.Error)
	}

	// Check if any rows were affected
	if result.RowsAffected == 0 {
		return fmt.Errorf("no user found with id: %s", user.ID)
	}

	return nil
}
