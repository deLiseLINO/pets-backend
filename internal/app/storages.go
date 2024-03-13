package app

import (
	"pets-backend/internal/ent"
	"pets-backend/internal/storage/postgres"
)

func (app *App) initOtpStorage(client *ent.Client) *postgres.OtpStorage {
	return postgres.NewOtpStorage(client)
}

func (app *App) initUserStorage(client *ent.Client) *postgres.UserStorage {
	return postgres.NewUserStorage(client)
}
