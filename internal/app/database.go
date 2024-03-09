package app

import (
	"context"
	"log"
	"pets-backend/internal/database"
	"pets-backend/internal/ent"
)

func (app *App) initDatabase() *ent.Client {
	db, err := database.Connect(&app.cfg.Database)
	if err != nil {
		log.Fatal(err)
	}

	// Auto Migration
	if err = database.CreateSchema(context.Background(), db); err != nil {
		log.Fatal(err)
	}

	return db
}
