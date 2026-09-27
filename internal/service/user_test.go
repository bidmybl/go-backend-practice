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
