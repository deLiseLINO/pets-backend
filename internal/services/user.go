package services

import (
	"context"
	"pets-backend/internal/models"
)

type UserStorate interface {
	GetByEmail(ctx context.Context, email string) (*models.User, error)
}

type UserService struct {
	UserStorage UserStorate
}

func NewUserService(userStorage UserStorate) *UserService {
	return &UserService{
		UserStorage: userStorage,
	}
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.UserStorage.GetByEmail(ctx, email)
}
