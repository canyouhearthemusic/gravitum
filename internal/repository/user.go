package repository

import (
	"context"
	"errors"
	"fmt"

	"github.com/canyouhearthemusic/gravitum/internal/entity"
	"github.com/canyouhearthemusic/gravitum/pkg/postgres"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
)

type UserRepository struct {
	db *postgres.Postgres
}

func NewUserRepository(db *postgres.Postgres) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *entity.User) error {
	sql, args, err := r.db.Builder.
		Insert("users").
		Columns("id", "username", "email", "first_name", "last_name", "created_at", "updated_at").
		Values(user.ID, user.Username, user.Email, user.FirstName, user.LastName, user.CreatedAt, user.UpdatedAt).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL: %w", err)
	}

	_, err = r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	return nil
}

func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*entity.User, error) {
	sql, args, err := r.db.Builder.
		Select("id", "username", "email", "first_name", "last_name", "created_at", "updated_at").
		From("users").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL: %w", err)
	}

	user := new(entity.User)
	err = r.db.Pool.QueryRow(ctx, sql, args...).Scan(
		&user.ID,
		&user.Username,
		&user.Email,
		&user.FirstName,
		&user.LastName,
		&user.CreatedAt,
		&user.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, fmt.Errorf("user not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get user: %w", err)
	}

	return user, nil
}

func (r *UserRepository) GetAll(ctx context.Context) ([]*entity.User, error) {
	sql, args, err := r.db.Builder.
		Select("id", "username", "email", "first_name", "last_name", "created_at", "updated_at").
		From("users").
		OrderBy("created_at DESC").
		ToSql()
	if err != nil {
		return nil, fmt.Errorf("failed to build SQL: %w", err)
	}

	rows, err := r.db.Pool.Query(ctx, sql, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to query users: %w", err)
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user entity.User
		err = rows.Scan(
			&user.ID,
			&user.Username,
			&user.Email,
			&user.FirstName,
			&user.LastName,
			&user.CreatedAt,
			&user.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("failed to scan user: %w", err)
		}
		users = append(users, &user)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error iterating users: %w", err)
	}

	return users, nil
}

func (r *UserRepository) Update(ctx context.Context, user *entity.User) error {
	sql, args, err := r.db.Builder.
		Update("users").
		Set("username", user.Username).
		Set("email", user.Email).
		Set("first_name", user.FirstName).
		Set("last_name", user.LastName).
		Set("updated_at", user.UpdatedAt).
		Where("id = ?", user.ID).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL: %w", err)
	}

	result, err := r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}

	if result.RowsAffected() == 0 {
		return errors.New("user not found")
	}

	return nil
}

func (r *UserRepository) Delete(ctx context.Context, id uuid.UUID) error {
	sql, args, err := r.db.Builder.
		Delete("users").
		Where("id = ?", id).
		ToSql()
	if err != nil {
		return fmt.Errorf("failed to build SQL: %w", err)
	}

	result, err := r.db.Pool.Exec(ctx, sql, args...)
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}

	if result.RowsAffected() == 0 {
		return errors.New("user not found")
	}

	return nil
}
