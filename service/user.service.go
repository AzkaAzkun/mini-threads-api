package service

import (
	"context"
	"errors"

	"github.com/AzkaAzkun/mini-threads-api/dto"
	"github.com/AzkaAzkun/mini-threads-api/entity"
	"github.com/AzkaAzkun/mini-threads-api/repository"
	"github.com/AzkaAzkun/mini-threads-api/utils"
	"gorm.io/gorm"
)

type IUserService interface {
	RegisterAccount(ctx context.Context, user dto.UserCreate) (string, error)
	Login(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error)
}

type UserService struct {
	userRepository repository.IUserRepository
	db             *gorm.DB
}

func NewUser(userRepository repository.IUserRepository,
	db *gorm.DB) IUserService {
	return &UserService{
		userRepository: userRepository,
		db:             db,
	}
}

func (s *UserService) RegisterAccount(ctx context.Context, user dto.UserCreate) (string, error) {
	createResult, err := s.userRepository.Create(ctx, nil, entity.User{
		Email:    user.Email,
		Name:     user.Name,
		Password: user.Password,
	})
	if err != nil {
		return "", err
	}

	return createResult.ID.String(), nil
}

func (s *UserService) Login(ctx context.Context, req dto.UserLoginRequest) (dto.UserLoginResponse, error) {
	user, err := s.userRepository.GetByEmail(ctx, nil, req.Email)
	if err != nil {
		return dto.UserLoginResponse{}, err
	}

	checkPassword, err := utils.CheckPassword(user.Password, []byte(req.Password))
	if err != nil || !checkPassword {
		return dto.UserLoginResponse{}, errors.New("invalid credentials")
	}

	payload := map[string]string{
		"user_id": user.ID.String(),
		"email":   user.Email,
	}

	token, err := utils.GenerateToken(payload, 24)
	if err != nil {
		return dto.UserLoginResponse{}, errors.New("something went wrong")
	}

	return dto.UserLoginResponse{
		Token: token,
	}, nil
}
