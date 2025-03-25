package user

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Model struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func New(dto *CreateDTO) (*Model, error) {
	user := &Model{
		Username:  dto.Username,
		Email:     dto.Email,
		FirstName: dto.FirstName,
		LastName:  dto.LastName}

	if err := user.validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *Model) Update(dto *UpdateDTO) error {
	u.Username = dto.Username
	u.Email = dto.Email
	u.FirstName = dto.FirstName
	u.LastName = dto.LastName
	u.UpdatedAt = time.Now()

	if err := u.validate(); err != nil {
		return err
	}

	return nil
}

func (u *Model) validate() error {
	if u.Username == "" {
		return errors.New("username is required")
	}

	if u.Email == "" {
		return errors.New("email is required")
	}

	if len(u.Username) < 3 {
		return errors.New("username must be at least 3 characters")
	}

	return nil
}
