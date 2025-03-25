package service

import (
	"context"
	"fmt"
	"log"

	"github.com/AzkaAzkun/mini-threads-api/dto"
	"github.com/AzkaAzkun/mini-threads-api/entity"
	"github.com/AzkaAzkun/mini-threads-api/repository"
	"github.com/AzkaAzkun/mini-threads-api/utils"
	"github.com/google/uuid"
	"github.com/oklog/ulid/v2"
	"gorm.io/gorm"
)

type IPostImageService interface {
	Create(ctx context.Context, post dto.PostImageCreate) (string, error)
	Delete(ctx context.Context, postId string) error
}

type PostImageService struct {
	postRepository repository.IPostImageRepository
	db             *gorm.DB
}

func NewPostImage(postRepository repository.IPostImageRepository,
	db *gorm.DB) IPostImageService {
	return &PostImageService{
		postRepository: postRepository,
		db:             db,
	}
}

func (s *PostImageService) Create(ctx context.Context, post dto.PostImageCreate) (string, error) {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Panic(r)
		}
	}()

	imagePath := fmt.Sprintf("post-image-%s.%s", ulid.Make(), utils.GetExtensions(post.Image.Filename))
	if err := utils.UploadFile(post.Image, imagePath); err != nil {
		tx.Rollback()
		return "", err
	}

	createResult, err := s.postRepository.Create(ctx, tx, entity.PostImage{
		PostId:    uuid.MustParse(post.PostId),
		ImagePath: imagePath,
	})
	if err != nil {
		tx.Rollback()
		return "", err
	}

	if err := tx.Commit().Error; err != nil {
		return "", err
	}

	return createResult.ID.String(), nil
}

func (s *PostImageService) Delete(ctx context.Context, postId string) error {
	post, err := s.postRepository.GetById(ctx, nil, postId)
	if err != nil {
		return err
	}

	return s.postRepository.Delete(ctx, nil, post)
}
