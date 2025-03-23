package service

import (
	"github.com/canyouhearthemusic/gravitum/internal/domain/user"
	"github.com/canyouhearthemusic/gravitum/internal/repository"
)

type Services struct {
	User user.Service
}

func New(repos *repository.Repositories) *Services {
	return &Services{
		User: NewUserService(repos.User),
	}
}
