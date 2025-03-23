package service

import (
	"context"
	"fmt"

	"github.com/canyouhearthemusic/gravitum/internal/entity"
	"github.com/canyouhearthemusic/gravitum/internal/repository"
	"github.com/google/uuid"
)

type UserService struct {
	repo *repository.UserRepository
}

func NewUserService(repo *repository.UserRepository) *UserService {
	return &UserService{repo: repo}
}

func (s *UserService) CreateUser(ctx context.Context, username, email, firstName, lastName string) (*entity.User, error) {
	user, err := entity.NewUser(username, email, firstName, lastName)
	if err != nil {
		return nil, fmt.Errorf("failed to create user entity: %w", err)
	}

	err = s.repo.Create(ctx, user)
	if err != nil {
		return nil, fmt.Errorf("failed to save user: %w", err)
	}

	return user, nil
}

func (s *UserService) GetUser(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	user, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]*entity.User, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to get all users: %w", err)
	}

	return users, nil
}

func (s *UserService) UpdateUser(ctx context.Context, id uuid.UUID, username, email, firstName, lastName string) (*entity.User, error) {
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
