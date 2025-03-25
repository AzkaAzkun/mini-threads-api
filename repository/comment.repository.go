package repository

import (
	"context"

	"github.com/AzkaAzkun/mini-threads-api/entity"
	"gorm.io/gorm"
)

type ICommentRepository interface {
	Create(ctx context.Context, tx *gorm.DB, comment entity.Comment) (entity.Comment, error)
	GetAllByPostId(ctx context.Context, tx *gorm.DB, postId string) ([]entity.Comment, error)
	GetById(ctx context.Context, tx *gorm.DB, commentId string) (entity.Comment, error)
	Update(ctx context.Context, tx *gorm.DB, commentId entity.Comment) error
	Delete(ctx context.Context, tx *gorm.DB, commentId entity.Comment) error
}

type CommentRepository struct {
	db *gorm.DB
}

func NewComment(db *gorm.DB) ICommentRepository {
	return &CommentRepository{db}
}

func (r *CommentRepository) Create(ctx context.Context, tx *gorm.DB, comment entity.Comment) (entity.Comment, error) {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Create(&comment).Error; err != nil {
		return comment, err
	}

	return comment, nil
}

func (r *CommentRepository) GetAllByPostId(ctx context.Context, tx *gorm.DB, postId string) ([]entity.Comment, error) {
	if tx == nil {
		tx = r.db
	}

	var comment []entity.Comment
	if err := tx.WithContext(ctx).Find(&comment, "post_id = ?", postId).Error; err != nil {
		return []entity.Comment{}, err
	}

	return comment, nil
}

func (r *CommentRepository) GetById(ctx context.Context, tx *gorm.DB, commentId string) (entity.Comment, error) {
	if tx == nil {
		tx = r.db
	}

	var comment entity.Comment
	if err := tx.WithContext(ctx).Take(&comment, "id = ?", commentId).Error; err != nil {
		return entity.Comment{}, err
	}

	return comment, nil
}

func (r *CommentRepository) Update(ctx context.Context, tx *gorm.DB, comment entity.Comment) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Save(&comment).Error; err != nil {
		return err
	}

	return nil
}

func (r *CommentRepository) Delete(ctx context.Context, tx *gorm.DB, comment entity.Comment) error {
	if tx == nil {
		tx = r.db
	}

	if err := tx.WithContext(ctx).Delete(&comment).Error; err != nil {
		return err
	}

	return nil
}
