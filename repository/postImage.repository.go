package repository

import (
	"context"

	"github.com/AzkaAzkun/mini-threads-api/entity"
	"gorm.io/gorm"
)

type IPostImageRepository interface {
	Create(ctx context.Context, tx *gorm.DB, postImage entity.PostImage) (entity.PostImage, error)
	GetById(ctx context.Context, tx *gorm.DB, postId string) (entity.PostImage, error)
	Delete(ctx context.Context, tx *gorm.DB, postImage entity.PostImage) error
}

type PostImageRepository struct {
	db *gorm.DB
}

func NewPostImage(db *gorm.DB) IPostImageRepository {
	return &PostImageRepository{db}
}

func (r *PostImageRepository) Create(ctx context.Context, tx *gorm.DB, postImage entity.PostImage) (entity.PostImage, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&postImage).Error; err != nil {
		return postImage, err
	}

	return postImage, nil
}

func (r *PostImageRepository) GetById(ctx context.Context, tx *gorm.DB, postImageId string) (entity.PostImage, error) {
	if tx == nil {
		tx = r.db
	}

	var postImage entity.PostImage
	if err := tx.WithContext(ctx).Take(&postImage, "id = ?", postImageId).Error; err != nil {
		return entity.PostImage{}, err
	}

	return postImage, nil
}

func (r *PostImageRepository) Delete(ctx context.Context, tx *gorm.DB, postImage entity.PostImage) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&postImage).Error; err != nil {
		return err
	}

	return nil
}
