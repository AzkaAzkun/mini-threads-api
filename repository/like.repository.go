package repository

import (
	"context"

	"github.com/AzkaAzkun/mini-threads-api/entity"
	"gorm.io/gorm"
)

type ILikeRepository interface {
	Create(ctx context.Context, tx *gorm.DB, like entity.Like) (entity.Like, error)
	GetById(ctx context.Context, tx *gorm.DB, likeId string) (entity.Like, error)
	GetAll(ctx context.Context, tx *gorm.DB) ([]entity.Like, error)
	CountLikeByPost(ctx context.Context, tx *gorm.DB, postId string) (int, error)
	Delete(ctx context.Context, tx *gorm.DB, likeId entity.Like) error
}

type LikeRepository struct {
	db *gorm.DB
}

func NewLike(db *gorm.DB) ILikeRepository {
	return &LikeRepository{db}
}

func (r *LikeRepository) Create(ctx context.Context, tx *gorm.DB, like entity.Like) (entity.Like, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&like).Error; err != nil {
		return like, err
	}

	return like, nil
}

func (r *LikeRepository) GetById(ctx context.Context, tx *gorm.DB, likeId string) (entity.Like, error) {
	if tx == nil {
		tx = r.db
	}

	var like entity.Like
	if err := tx.WithContext(ctx).Take(&like, "id = ?", likeId).Error; err != nil {
		return entity.Like{}, err
	}

	return like, nil
}

func (r *LikeRepository) GetAll(ctx context.Context, tx *gorm.DB) ([]entity.Like, error) {
	if tx == nil {
		tx = r.db
	}

	var like []entity.Like
	if err := tx.WithContext(ctx).Preload("Comment").Preload("Like").Find(&like).Error; err != nil {
		return []entity.Like{}, err
	}

	return like, nil
}

func (r *LikeRepository) CountLikeByPost(ctx context.Context, tx *gorm.DB, postId string) (int, error) {
	if tx == nil {
		tx = r.db
	}

	var count int64
	if err := tx.WithContext(ctx).Model(&entity.Like{}).Where("post_id = ?", postId).Count(&count).Error; err != nil {
		return 0, err
	}

	return int(count), nil
}

func (r *LikeRepository) Delete(ctx context.Context, tx *gorm.DB, like entity.Like) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&like).Error; err != nil {
		return err
	}

	return nil
}
