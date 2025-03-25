package service

import (
	"context"

	"github.com/AzkaAzkun/mini-threads-api/dto"
	"github.com/AzkaAzkun/mini-threads-api/entity"
	"github.com/AzkaAzkun/mini-threads-api/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ILikeService interface {
	Create(ctx context.Context, like dto.LikeCreate) (string, error)
	GetAll(ctx context.Context) ([]entity.Like, error)
	Delete(ctx context.Context, likeId string) error
}

type LikeService struct {
	likeRepository repository.ILikeRepository
	postRepository repository.IPostRepository
	db             *gorm.DB
}

func NewLike(likeRepository repository.ILikeRepository,
	postRepository repository.IPostRepository,
	db *gorm.DB) ILikeService {
	return &LikeService{
		likeRepository: likeRepository,
		postRepository: postRepository,
		db:             db,
	}
}

func (s *LikeService) Create(ctx context.Context, like dto.LikeCreate) (string, error) {
	post, err := s.postRepository.GetById(ctx, nil, like.PostId)
	if err != nil {
		return "", err
	}

	createResult, err := s.likeRepository.Create(ctx, nil, entity.Like{
		UserId: uuid.MustParse(like.UserId),
		PostId: post.ID,
	})
	if err != nil {
		return "", err
	}

	countLike, err := s.likeRepository.CountLikeByPost(ctx, nil, like.PostId)
	if err != nil {
		return "", err
	}

	post.LikeCount = countLike

	if err := s.postRepository.Update(ctx, nil, post); err != nil {
		return "", err
	}

	return createResult.ID.String(), nil
}

func (s *LikeService) GetAll(ctx context.Context) ([]entity.Like, error) {
	likes, err := s.likeRepository.GetAll(ctx, nil)
	if err != nil {
		return []entity.Like{}, err
	}

	return likes, nil
}

func (s *LikeService) Delete(ctx context.Context, likeId string) error {
	like, err := s.likeRepository.GetById(ctx, nil, likeId)
	if err != nil {
		return err
	}

	if err := s.likeRepository.Delete(ctx, nil, like); err != nil {
		return err
	}

	post, err := s.postRepository.GetById(ctx, nil, like.PostId.String())
	if err != nil {
		return err
	}

	countLike, err := s.likeRepository.CountLikeByPost(ctx, nil, post.ID.String())
	if err != nil {
		return err
	}

	post.LikeCount = countLike

	if err := s.postRepository.Update(ctx, nil, post); err != nil {
		return err
	}

	return nil
}
