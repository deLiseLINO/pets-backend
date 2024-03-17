package postgres

import (
	"context"
	"pets-backend/internal/ent"
	"pets-backend/internal/ent/user"
	"pets-backend/internal/models"

	"github.com/google/uuid"
)

type UserStorage struct {
	client *ent.Client
}

func NewUserStorage(client *ent.Client) *UserStorage {
	return &UserStorage{client: client}
}

func (s *UserStorage) Add(
	ctx context.Context,
	username string,
	name string,
	email string,
) (*models.User, error) {
	usr, err := s.client.User.
		Create().
		SetID(uuid.New()).
		SetUniqueName(username).
		SetName(name).
		SetEmail(email).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return userToModel(usr), nil
}

func (s *UserStorage) AddWithPassword(
	ctx context.Context,
	username string,
	name string,
	email string,
	password string,
) (*models.User, error) {
	usr, err := s.client.User.
		Create().
		SetID(uuid.New()).
		SetUniqueName(username).
		SetName(name).
		SetEmail(email).
		SetPassword(password).
		Save(ctx)
	if err != nil {
		return nil, err
	}
	return userToModel(usr), err
}

func (s *UserStorage) GetByEmail(
	ctx context.Context,
	email string,
) (*models.User, error) {
	usr, err := s.client.User.
		Query().
		Where(user.Email(email)).
		First(ctx)
	if err != nil {
		switch {
		case ent.IsNotFound(err):
			return nil, models.ErrUserNotFound
		default:
			return nil, err
		}
	}
	return userToModel(usr), nil
}

func userToModel(user *ent.User) *models.User {
	return &models.User{
		ID:         user.ID,
		Email:      user.Email,
		UniqueName: user.UniqueName,
		Name:       user.Name,
	}
}
