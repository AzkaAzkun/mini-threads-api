package repository

import (
	"context"

	"github.com/AzkaAzkun/mini-threads-api/entity"
	"gorm.io/gorm"
)

type IPostRepository interface {
	Create(ctx context.Context, tx *gorm.DB, post entity.Post) (entity.Post, error)
	GetById(ctx context.Context, tx *gorm.DB, postId string) (entity.Post, error)
	GetAll(ctx context.Context, tx *gorm.DB) ([]entity.Post, error)
	Update(ctx context.Context, tx *gorm.DB, post entity.Post) error
	Delete(ctx context.Context, tx *gorm.DB, post entity.Post) error
}

type PostRepository struct {
	db *gorm.DB
}

func NewPost(db *gorm.DB) IPostRepository {
	return &PostRepository{db}
}

func (r *PostRepository) Create(ctx context.Context, tx *gorm.DB, post entity.Post) (entity.Post, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&post).Error; err != nil {
		return post, err
	}

	return post, nil
}

func (r *PostRepository) GetById(ctx context.Context, tx *gorm.DB, postId string) (entity.Post, error) {
	if tx == nil {
		tx = r.db
	}

	var post entity.Post
	if err := tx.WithContext(ctx).Preload("Comment").Preload("Like").Preload("PostImage").
		Take(&post, "id = ?", postId).Error; err != nil {
		return entity.Post{}, err
	}

	return post, nil
}

func (r *PostRepository) GetAll(ctx context.Context, tx *gorm.DB) ([]entity.Post, error) {
	if tx == nil {
		tx = r.db
	}

	var post []entity.Post
	if err := tx.WithContext(ctx).Preload("Like").Preload("PostImage").
		Find(&post).Error; err != nil {
		return []entity.Post{}, err
	}

	return post, nil
}

func (r *PostRepository) Update(ctx context.Context, tx *gorm.DB, post entity.Post) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Save(&post).Error; err != nil {
		return err
	}

	return nil
}

func (r *PostRepository) Delete(ctx context.Context, tx *gorm.DB, post entity.Post) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&post).Error; err != nil {
		return err
	}

	return nil
}
