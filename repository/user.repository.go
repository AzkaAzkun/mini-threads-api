package repository

import (
	"context"

	"github.com/AzkaAzkun/mini-threads-api/entity"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Create(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error)
	GetById(ctx context.Context, tx *gorm.DB, userId string) (entity.User, error)
	GetByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) IUserRepository {
	return &UserRepository{db}
}

func (r *UserRepository) Create(ctx context.Context, tx *gorm.DB, user entity.User) (entity.User, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *UserRepository) GetById(ctx context.Context, tx *gorm.DB, userId string) (entity.User, error) {
	if tx == nil {
		tx = r.db
	}

	var user entity.User
	if err := tx.WithContext(ctx).Take(&user, "id = ?", userId).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r *UserRepository) GetByEmail(ctx context.Context, tx *gorm.DB, email string) (entity.User, error) {
	if tx == nil {
		tx = r.db
	}

	var user entity.User
	if err := tx.WithContext(ctx).Take(&user, "email = ?", email).Error; err != nil {
		return entity.User{}, err
	}

	return user, nil
}
