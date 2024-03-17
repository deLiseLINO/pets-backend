package services

import (
	"context"
	"pets-backend/internal/models"
)

type UserStorate interface {
	GetByEmail(ctx context.Context, email string) (*models.User, error)
	Add(ctx context.Context, username string, name string, email string) (*models.User, error)
	AddWithPassword(ctx context.Context, username string, name string, email string, password string) (*models.User, error)
}

type UserService struct {
	UserStorage UserStorate
}

func NewUserService(userStorage UserStorate) *UserService {
	return &UserService{
		UserStorage: userStorage,
	}
}

func (s *UserService) Add(ctx context.Context, username string, name string, email string) (*models.User, error) {
	return s.UserStorage.Add(ctx, username, name, email)
}

func (s *UserService) AddWithPassword(
	ctx context.Context,
	username string,
	name string,
	email string,
	password string,
) (*models.User, error) {
	return s.UserStorage.AddWithPassword(ctx, username, name, email, password)
}

func (s *UserService) GetByEmail(ctx context.Context, email string) (*models.User, error) {
	return s.UserStorage.GetByEmail(ctx, email)
}
