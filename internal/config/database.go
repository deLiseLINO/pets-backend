package config

import (
	"fmt"
	"strings"
)

type Database struct {
	Address  string `yaml:"address" validate:"required"`
	Username string `yaml:"username" validate:"required"`
	Password string `yaml:"password" validate:"required"`
	Database string `yaml:"database" validate:"required"`
}

const dsnLength = 2

func (cfg Database) BuildDSN() string {
	data := strings.Split(cfg.Address, ":")
	if len(data) != dsnLength {
		panic("address must be in the host:port format")
	}

	return fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		data[0],
		data[1],
		cfg.Username,
		cfg.Password,
		cfg.Database,
	)
}
