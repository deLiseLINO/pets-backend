package main

import (
	"log"
	"pets-backend/internal/app"
	"pets-backend/internal/config"

	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatal(err)
	}

	application := app.New(cfg)
	application.Run()
}
