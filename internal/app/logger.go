package app

import (
	"os"
	"pets-backend/internal/config"

	log "github.com/sirupsen/logrus"
)

func (app *App) initLogger(loggerCfg *config.Logger) {
	log.SetFormatter(&log.JSONFormatter{
		PrettyPrint: true,
	})
	log.SetReportCaller(true)
	log.SetLevel(log.Level(loggerCfg.Level))

	switch loggerCfg.Output {
	case config.OutputStdout:
		log.SetOutput(os.Stdout)

	case config.OutputStderr:
		log.SetOutput(os.Stderr)

	default:
		panic("unsupported logger output")
	}
}
