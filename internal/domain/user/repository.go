package user

import (
	"context"

	"github.com/google/uuid"
)

type Repository interface {
	Create(ctx context.Context, user *Model) error
	GetByID(ctx context.Context, id uuid.UUID) (*Model, error)
	GetAll(ctx context.Context) ([]*Model, error)
	Update(ctx context.Context, user *Model) error
	Delete(ctx context.Context, id uuid.UUID) error
}
