package user

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	tests := []struct {
		name      string
		username  string
		email     string
		firstName string
		lastName  string
		wantErr   bool
	}{
		{
			name:      "Valid user",
			username:  "testuser",
			email:     "test@example.com",
			firstName: "Test",
			lastName:  "User",
			wantErr:   false,
		},
		{
			name:      "Empty username",
			username:  "",
			email:     "test@example.com",
			firstName: "Test",
			lastName:  "User",
			wantErr:   true,
		},
		{
			name:      "Short username",
			username:  "te",
			email:     "test@example.com",
			firstName: "Test",
			lastName:  "User",
			wantErr:   true,
		},
		{
			name:      "Empty email",
			username:  "testuser",
			email:     "",
			firstName: "Test",
			lastName:  "User",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dto := &CreateDTO{
				Username:  tt.username,
				Email:     tt.email,
				FirstName: tt.firstName,
				LastName:  tt.lastName,
			}
			user, err := New(dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewUser() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil {
				if user.Username != tt.username {
					t.Errorf("NewUser() username = %v, want %v", user.Username, tt.username)
				}
				if user.Email != tt.email {
					t.Errorf("NewUser() email = %v, want %v", user.Email, tt.email)
				}
				if user.FirstName != tt.firstName {
					t.Errorf("NewUser() firstName = %v, want %v", user.FirstName, tt.firstName)
				}
				if user.LastName != tt.lastName {
					t.Errorf("NewUser() lastName = %v, want %v", user.LastName, tt.lastName)
				}
				if user.ID.String() == "" {
					t.Errorf("NewUser() id should not be empty")
				}
			}
		})
	}
}

func TestUser_Update(t *testing.T) {
	dto := &CreateDTO{
		Username:  "olduser",
		Email:     "old@example.com",
		FirstName: "Old",
		LastName:  "User",
	}
	user, _ := New(dto)

	tests := []struct {
		name       string
		username   string
		email      string
		firstName  string
		lastName   string
		wantErr    bool
		wantResult *Model
	}{
		{
			name:      "Update all fields",
			username:  "newuser",
			email:     "new@example.com",
			firstName: "New",
			lastName:  "User",
			wantErr:   false,
			wantResult: &Model{
				ID:        user.ID,
				Username:  "newuser",
				Email:     "new@example.com",
				FirstName: "New",
				LastName:  "User",
			},
		},
		{
			name:      "Update partial fields",
			username:  "",
			email:     "",
			firstName: "Updated",
			lastName:  "",
			wantErr:   false,
			wantResult: &Model{
				ID:        user.ID,
				Username:  "newuser",
				Email:     "new@example.com",
				FirstName: "Updated",
				LastName:  "User",
			},
		},
		{
			name:      "Invalid username",
			username:  "ab",
			email:     "",
			firstName: "",
			lastName:  "",
			wantErr:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dto := &UpdateDTO{
				Username:  tt.username,
				Email:     tt.email,
				FirstName: tt.firstName,
				LastName:  tt.lastName,
			}
			err := user.Update(dto)
			if (err != nil) != tt.wantErr {
				t.Errorf("User.Update() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if err == nil && tt.wantResult != nil {
				if user.Username != tt.wantResult.Username {
					t.Errorf("User.Update() username = %v, want %v", user.Username, tt.wantResult.Username)
				}
				if user.Email != tt.wantResult.Email {
					t.Errorf("User.Update() email = %v, want %v", user.Email, tt.wantResult.Email)
				}
				if user.FirstName != tt.wantResult.FirstName {
					t.Errorf("User.Update() firstName = %v, want %v", user.FirstName, tt.wantResult.FirstName)
				}
				if user.LastName != tt.wantResult.LastName {
					t.Errorf("User.Update() lastName = %v, want %v", user.LastName, tt.wantResult.LastName)
				}
				if user.UpdatedAt.IsZero() {
					t.Errorf("User.Update() UpdatedAt should not be zero")
				}
			}
		})
	}
}
