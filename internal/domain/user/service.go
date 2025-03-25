package user

import (
	"context"

	"github.com/google/uuid"
)

type Service interface {
	CreateUser(ctx context.Context, dto *CreateDTO) error
	GetUser(ctx context.Context, id uuid.UUID) (*Model, error)
	GetAllUsers(ctx context.Context) ([]*Model, error)
	UpdateUser(ctx context.Context, id uuid.UUID, dto *UpdateDTO) error
	DeleteUser(ctx context.Context, id uuid.UUID) error
}
