package main

import (
	"log"

	"github.com/evrone/go-clean-template/internal/app"

	"github.com/evrone/go-clean-template/config"
)

func main() {
	cfg, err := config.NewViperConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
