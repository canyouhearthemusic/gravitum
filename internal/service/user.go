package service

import (
	"context"
	"fmt"

	"github.com/canyouhearthemusic/gravitum/internal/domain/user"
	"github.com/google/uuid"
)

type UserService struct {
	repo user.Repository
}

func NewUserService(repo user.Repository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, username, email, firstName, lastName string) (*user.Model, error) {
	user, err := user.New(username, email, firstName, lastName)
	if err != nil {
		return nil, fmt.Errorf("failed to create user entity: %w", err)
	}

	err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return user, nil
}

func (s *UserService) GetUser(ctx context.Context, id uuid.UUID) (*user.Model, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*user.Model, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}

	return users, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id uuid.UUID, username, email, firstName, lastName string) (*user.Model, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to find user: %w", err)
	}

	err = user.Update(username, email, firstName, lastName)
	if err != nil {
		return nil, fmt.Errorf("failed to update user data: %w", err)
	}

	err = s.repo.Update(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to save user changes: %w", err)
	}

	return user, nil
}

func (s *UserService) DeleteUser(ctx context.Context, id uuid.UUID) error {
	err := s.repo.Delete(ctx, id)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	return nil
}
