package app

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/canyouhearthemusic/gravitum/internal/config"
	"github.com/canyouhearthemusic/gravitum/internal/handler"
	"github.com/canyouhearthemusic/gravitum/internal/repository"
	"github.com/canyouhearthemusic/gravitum/internal/service"
	"github.com/canyouhearthemusic/gravitum/pkg/httpserver"
	"github.com/canyouhearthemusic/gravitum/pkg/logger"
	"github.com/canyouhearthemusic/gravitum/pkg/postgres"
)

func Run(cfg *config.Config) {
	l := logger.New(cfg.Log.Level)

	pg, err := postgres.New(cfg.Database)
	if err != nil {
		l.Fatal("Postgres error: %s", err)
	}
	defer pg.Close()

	repositories := repository.New(pg)

	services := service.New(repositories)

	server := httpserver.New(httpserver.Port(cfg.App.Port))

	handlers := handler.New(services)
	handlers.Register(server.App)

	server.Start()
	l.Info("Server started on port %s", cfg.App.Port)

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGTERM)

	select {
	case s := <-interrupt:
		l.Info("app - Run - signal: " + s.String())
	case err = <-server.Notify():
		l.Error(fmt.Errorf("app - Run - server.Notify: %w", err))
	}

	err = server.Shutdown()
	if err != nil {
		l.Error(fmt.Errorf("app - Run - server.Shutdown: %w", err))
	}

	l.Info("app - Run - server gracefully stopped")
}
