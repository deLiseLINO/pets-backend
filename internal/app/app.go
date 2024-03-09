package app

import "pets-backend/internal/config"

type App struct {
	cfg *config.Config
}

func New(cfg *config.Config) *App {
	return &App{cfg: cfg}
}

func (app *App) Run() {
	app.initLogger(&app.cfg.Logger)
	_ = app.initDatabase()
	router := app.initRouter()
	app.runServer(router)
}
