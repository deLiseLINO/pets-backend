package postgres

import (
	"context"
	"pets-backend/internal/ent"
)

type UserStorage struct {
	client *ent.Client
}

func NewUserStorage(client *ent.Client) *UserStorage {
	return &UserStorage{client: client}
}

func (s *UserStorage) GetByEmail(
	ctx context.Context,
	email string,
) error {
	return nil
}
