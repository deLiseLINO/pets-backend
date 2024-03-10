package app

import (
	"pets-backend/internal/config"
)

type App struct {
	cfg *config.Config
}

func New(cfg *config.Config) *App {
	return &App{cfg: cfg}
}

func (app *App) Run() {
	app.initLogger(&app.cfg.Logger)
	connection := app.initDatabase()

	otpStorage := app.initOtpStorage(connection)
	otpSvc := app.initOtpSender(app.cfg.SMTP, otpStorage)

	router := app.initRouter(otpSvc)
	app.runServer(router)
}
