package repository

import (
	"github.com/canyouhearthemusic/gravitum/internal/domain/user"
)

type Repositories struct {
	User user.Repository
}

func New(cfgs ...Configuration) (*Repositories, error) {
	repos := new(Repositories)

	for _, cfg := range cfgs {
		if err := cfg(repos); err != nil {
			return nil, err
		}
	}

	return repos, nil
}
