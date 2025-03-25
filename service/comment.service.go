package service

import (
	"context"

	"github.com/AzkaAzkun/mini-threads-api/dto"
	"github.com/AzkaAzkun/mini-threads-api/entity"
	"github.com/AzkaAzkun/mini-threads-api/repository"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ICommentService interface {
	Create(ctx context.Context, comment dto.CommentCreate) (string, error)
	GetAllByPostId(ctx context.Context, postId string) ([]entity.Comment, error)
	Edit(ctx context.Context, req dto.CommentUpdate) error
	Delete(ctx context.Context, commentId string) error
}

type CommentService struct {
	commentRepository repository.ICommentRepository
	postRepository    repository.IPostRepository
	db                *gorm.DB
}

func NewComment(commentRepository repository.ICommentRepository,
	postRepository repository.IPostRepository,
	db *gorm.DB) ICommentService {
	return &CommentService{
		commentRepository: commentRepository,
		postRepository:    postRepository,
		db:                db,
	}
}

func (s *CommentService) Create(ctx context.Context, comment dto.CommentCreate) (string, error) {
	post, err := s.postRepository.GetById(ctx, nil, comment.PostId)
	if err != nil {
		return "", err
	}

	createResult, err := s.commentRepository.Create(ctx, nil, entity.Comment{
		PostId: post.ID,
		Body:   comment.Body,
		UserId: uuid.MustParse(comment.UserId),
	})
	if err != nil {
		return "", err
	}

	return createResult.ID.String(), nil
}

func (s *CommentService) GetAllByPostId(ctx context.Context, postId string) ([]entity.Comment, error) {
	comments, err := s.commentRepository.GetAllByPostId(ctx, nil, postId)
	if err != nil {
		return []entity.Comment{}, err
	}

	return comments, nil
}

func (s *CommentService) Edit(ctx context.Context, req dto.CommentUpdate) error {
	comment, err := s.commentRepository.GetById(ctx, nil, req.CommentId)
	if err != nil {
		return err
	}

	comment.Body = req.Body

	if err := s.commentRepository.Update(ctx, nil, comment); err != nil {
		return err
	}

	return nil
}

func (s *CommentService) Delete(ctx context.Context, commentId string) error {
	comment, err := s.commentRepository.GetById(ctx, nil, commentId)
	if err != nil {
		return err
	}

	return s.commentRepository.Delete(ctx, nil, comment)
}
