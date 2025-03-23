package repository

import (
	pgRepo "github.com/canyouhearthemusic/gravitum/internal/repository/postgres"
	"github.com/canyouhearthemusic/gravitum/pkg/postgres"
)

type Configuration func(r *Repositories) error

func WithPostgresStore(pg *postgres.Postgres) Configuration {
	return func(r *Repositories) (err error) {
		r.User = pgRepo.NewUserRepository(pg)

		return nil
	}
}
