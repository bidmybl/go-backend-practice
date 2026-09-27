package service

import (
	"errors"
	"testing"

	"github.com/bidmybl/go-backend-practice/internal/model"
)

func TestGetUserByID(t *testing.T) {
	tests := []struct {
		name    string
		userID  int
		wantID  int
		wantErr error
	}{
		{
			name:   "user exists",
			userID: 1,
			wantID: 1,
		},
		{
			name:    "user not found",
			userID:  999,
			wantErr: ErrUserNotFound,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := &UserService{
				users: []model.User{
					{ID: 1, Name: "Denis", Age: 19},
					{ID: 2, Name: "Alice", Age: 20},
				},
			}

			user, err := service.GetUserByID(tt.userID)

			if tt.wantErr != nil {
				if !errors.Is(err, tt.wantErr) {
					t.Errorf("expected error %v, got %v", tt.wantErr, err)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: got %v", err)
			}

			if user.ID != tt.wantID {
				t.Errorf("expected user ID %d, got %d", tt.wantID, user.ID)
			}
		})
	}
}

func TestCreateUser(t *testing.T) {
	tests := []struct {
		name    string
		user    model.User
		wantErr error
	}{
		{
			name: "successful creation",
			user: model.User{
				Name: "Alice",
				Age:  23,
			},
		},
		{
			name: "empty name error",
			user: model.User{
				Name: "",
			},
			wantErr: ErrInvalidUser,
		},
		{
			name: "invalid age error",
			user: model.User{
				Age: 0,
			},
			wantErr: ErrInvalidUser,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			service := NewUserService()
			err := service.CreateUser(tt.user)

			if !errors.Is(err, tt.wantErr) {
				t.Fatalf("expected error %v, got %v", tt.wantErr, err)
			}
		})
	}
}

func TestGetUsersByName(t *testing.T) {
	tests := []struct {
		name       string
		searchName string
		wantFound  bool
	}{
		{
			name:       "user exists",
			searchName: "Alice",
			wantFound:  true,
		},
		{
			name:       "user does not exist",
			searchName: "Bob",
			wantFound:  false,
		},
	}

	service := &UserService{
		users: []model.User{{
			Name: "Alice",
			Age:  23},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			users := service.GetUsersByName(tt.searchName)
			if tt.wantFound && len(users) == 0 {
				t.Fatalf("expected to find user %s but got none", tt.searchName)
			}

			if !tt.wantFound && len(users) > 0 {
				t.Fatalf("expected no users for %s but got %d", tt.searchName, len(users))
			}
		})
	}
}

func TestGetUsers(t *testing.T) {
	service := NewUserService()
	users := service.GetUsers()
	if len(users) != 0 {
		t.Fatalf("expected 0 users, got %d", len(users))
	}

	_ = service.CreateUser(model.User{Name: "Alice", Age: 23})
	users = service.GetUsers()
	if len(users) != 1 {
		t.Fatalf("expected 1 user, got %d", len(users))
	}
}
