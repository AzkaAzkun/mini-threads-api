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

type IPostService interface {
	Create(ctx context.Context, post dto.PostCreate) (string, error)
	GetAll(ctx context.Context) ([]entity.Post, error)
	GetById(ctx context.Context, postId string) (entity.Post, error)
	Edit(ctx context.Context, req dto.PostUpdate) error
	Delete(ctx context.Context, postId string) error
}

type PostService struct {
	postRepository repository.IPostRepository
	db             *gorm.DB
}

func NewPost(postRepository repository.IPostRepository,
	db *gorm.DB) IPostService {
	return &PostService{
		postRepository: postRepository,
		db:             db,
	}
}

func (s *PostService) Create(ctx context.Context, post dto.PostCreate) (string, error) {
	tx := s.db.Begin()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
			log.Panicln("Rolled back due to panic: ", r)
		}
	}()

	postId := uuid.New()
	var postImage []entity.PostImage
	for _, image := range post.Image {
		imagePath := fmt.Sprintf("post-image-%s.%s", ulid.Make(), utils.GetExtensions(image.Filename))
		err := utils.UploadFile(image, imagePath)
		if err != nil {
			return "", err
		}

		postImage = append(postImage, entity.PostImage{
			PostId:    postId,
			ImagePath: imagePath,
		})
	}

	createResult, err := s.postRepository.Create(ctx, tx, entity.Post{
		UserId:    uuid.MustParse(post.UserId),
		Title:     post.Title,
		Body:      post.Body,
		PostImage: postImage,
	})
	if err != nil {
		tx.Rollback()
		return "", err
	}

	if err := tx.Commit().Error; err != nil {
		return "", err
	}

	log.Println("Created Post ID:", createResult.ID)
	log.Println("Inserted Images:", postImage)

	return createResult.ID.String(), nil
}

func (s *PostService) GetAll(ctx context.Context) ([]entity.Post, error) {
	posts, err := s.postRepository.GetAll(ctx, nil)
	if err != nil {
		return []entity.Post{}, err
	}

	return posts, nil
}

func (s *PostService) GetById(ctx context.Context, postId string) (entity.Post, error) {
	post, err := s.postRepository.GetById(ctx, nil, postId)
	if err != nil {
		return entity.Post{}, nil
	}

	return post, nil
}

func (s *PostService) Edit(ctx context.Context, req dto.PostUpdate) error {
	post, err := s.postRepository.GetById(ctx, nil, req.PostId)
	if err != nil {
		return err
	}

	post.Title = req.Title
	post.Body = req.Body

	if err := s.postRepository.Update(ctx, nil, post); err != nil {
		return err
	}

	return nil
}

func (s *PostService) Delete(ctx context.Context, postId string) error {
	post, err := s.postRepository.GetById(ctx, nil, postId)
	if err != nil {
		return err
	}

	return s.postRepository.Delete(ctx, nil, post)
}
