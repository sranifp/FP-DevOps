package service

import (
	"context"

	"FP-DevOps/dto"
	"FP-DevOps/entity"
	"FP-DevOps/repository"
	"FP-DevOps/utils"

	"gorm.io/gorm"
)

type (
	UserService interface {
		RegisterUser(ctx context.Context, req dto.UserRequest) (dto.UserResponse, error)
		Login(ctx context.Context, nrp string, password string) (entity.User, error)
		Me(ctx context.Context, userID string) (dto.UserResponse, error)
	}

	userService struct {
		userRepo repository.UserRepository
	}
)

func NewUserService(ur repository.UserRepository) UserService {
	return &userService{
		userRepo: ur,
	}
}

func (s *userService) RegisterUser(ctx context.Context, req dto.UserRequest) (dto.UserResponse, error) {
	_, err := s.userRepo.GetUserByUsername(req.Username)
	if err == nil && err != gorm.ErrRecordNotFound {
		return dto.UserResponse{}, dto.ErrUsernameAlreadyExists
	}

	user := entity.User{
		Username: req.Username,
		Password: req.Password,
	}

	userReg, err := s.userRepo.Create(user)
	if err != nil {
		return dto.UserResponse{}, dto.ErrCreateUser
	}

	return dto.UserResponse{
		ID:       userReg.ID.String(),
		Username: userReg.Username,
	}, nil
}

func (s *userService) Login(ctx context.Context, username string, password string) (entity.User, error) {
	user, err := s.userRepo.GetUserByUsername(username)
	if err != nil {
		return entity.User{}, dto.ErrCredentialsNotMatched
	}

	checkPassword, err := utils.CheckPassword(user.Password, []byte(password))
	if err != nil || !checkPassword {
		return entity.User{}, dto.ErrCredentialsNotMatched
	}

	return user, nil
}

func (s *userService) Me(ctx context.Context, userID string) (dto.UserResponse, error) {
	user, err := s.userRepo.GetUserById(userID)
	if err != nil {
		return dto.UserResponse{}, dto.ErrGetUserById
	}

	return dto.UserResponse{
		ID:       user.ID.String(),
		Username: user.Username,
	}, nil
}
