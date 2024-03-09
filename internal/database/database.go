package database

import (
	"context"
	"pets-backend/internal/config"
	"pets-backend/internal/ent"
)

func Connect(cfg *config.Database) (*ent.Client, error) {
	client, err := ent.Open("postgres", cfg.BuildDSN())
	if err != nil {
		return nil, err
	}

	return client, nil
}

// Auto migrations
func CreateSchema(ctx context.Context, client *ent.Client) error {
	return client.Schema.Create(ctx)
}
