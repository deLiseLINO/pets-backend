package config

import (
	"errors"
	"os"

	"github.com/andrew528i/yacl"
	"github.com/go-playground/validator/v10"
	"github.com/joho/godotenv"
)

type Config struct {
	Database   Database   `yaml:"database" validate:"required"`
	HTTPServer HTTPServer `yaml:"http_server" validate:"required"`
	Logger     Logger     `yaml:"logger" validate:"required"`
	SMTP       SMTP       `yaml:"smtp" validate:"required"`
	SSO        SSO        `yaml:"SSO" validate:"required"`
}

const (
	envPrefix   = "PETS_BACKEND"
	envFilePath = ".env"
)

func Parse() (*Config, error) {
	err := godotenv.Load(envFilePath)
	if err != nil && !errors.Is(err, os.ErrNotExist) {
		return nil, err
	}

	y := yacl.New[Config]()
	y.SetEnvPrefix(envPrefix)

	cfg, err := y.Parse()
	if err != nil {
		return nil, err
	}

	if err = validator.New().Struct(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
