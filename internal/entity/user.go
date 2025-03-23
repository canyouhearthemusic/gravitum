package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func NewUser(username, email, firstName, lastName string) (*User, error) {
	user := &User{
		Username:  username,
		Email:     email,
		FirstName: firstName,
		LastName:  lastName}

	if err := user.validate(); err != nil {
		return nil, err
	}

	return user, nil
}

func (u *User) Update(username, email, firstName, lastName string) error {
	u.Username = username
	u.Email = email
	u.FirstName = firstName
	u.LastName = lastName
	u.UpdatedAt = time.Now()

	if err := u.validate(); err != nil {
		return err
	}

	return nil
}

func (u *User) validate() error {
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
