package repository

import "github.com/canyouhearthemusic/gravitum/pkg/postgres"

type Repositories struct {
	User *UserRepository
}

func New(pg *postgres.Postgres) *Repositories {
	return &Repositories{
		User: NewUserRepository(pg),
	}
}
