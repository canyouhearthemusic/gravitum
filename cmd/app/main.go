package main

import (
	"log"

	"github.com/canyouhearthemusic/gravitum/internal/app"
	"github.com/canyouhearthemusic/gravitum/internal/config"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatalf("Config error: %s", err)
	}

	app.Run(cfg)
}
