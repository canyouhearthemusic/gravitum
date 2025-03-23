package service

import (
	"context"

	"github.com/canyouhearthemusic/gravitum/internal/entity"
	"github.com/canyouhearthemusic/gravitum/internal/repository"
	"github.com/google/uuid"
)

type UserServiceInterface interface {
	CreateUser(ctx context.Context, username, email, firstName, lastName string) (*entity.User, error)
	GetUser(ctx context.Context, id uuid.UUID) (*entity.User, error)
	GetAllUsers(ctx context.Context) ([]*entity.User, error)
	UpdateUser(ctx context.Context, id uuid.UUID, username, email, firstName, lastName string) (*entity.User, error)
	DeleteUser(ctx context.Context, id uuid.UUID) error
}

type Services struct {
	User UserServiceInterface
}

func New(repos *repository.Repositories) *Services {
	return &Services{
		User: NewUserService(repos.User),
	}
}
