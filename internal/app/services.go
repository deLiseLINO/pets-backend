package app

import (
	"pets-backend/internal/config"
	"pets-backend/internal/services"
	"pets-backend/internal/storage/postgres"
)

func (app *App) initOtpSender(cfg config.SMTP, otpStorage *postgres.OtpStorage) *services.OtpService {
	return services.NewOtpService(
		otpStorage,
		cfg.User,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.OtpLength,
		cfg.NextSendPeriodSec,
		cfg.ExparationTimeMin,
	)
}
