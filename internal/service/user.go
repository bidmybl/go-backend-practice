package service

import (
	"errors"
	"github.com/bidmybl/go-backend-practice/internal/model"
)

type UserService struct {
	users []model.User
}

var ErrInvalidUser = errors.New("invalid user")

func (s *UserService) CreateUser(user model.User) error {
	if user.Name == "" || user.Age <= 0 {
		return ErrInvalidUser
	}

	s.users = append(s.users, user)

	return nil
}

func (s *UserService) GetUsers(name string) []model.User {
	if name == "" {
		return s.users
	}

	var result []model.User

	for _, user := range s.users {
		if user.Name == name {
			result = append(result, user)
		}
	}

	return result
}
