package service

import (
	"errors"
	"github.com/bidmybl/go-backend-practice/internal/model"
)

type UserService struct {
	users  []model.User
	nextID int
}

func NewUserService() *UserService {
	return &UserService{
		nextID: 1,
	}
}

var ErrInvalidUser = errors.New("invalid user")
var ErrUserNotFound = errors.New("user not found")

func (s *UserService) CreateUser(user model.User) error {
	if user.Name == "" || user.Age <= 0 {
		return ErrInvalidUser
	}
	user.ID = s.nextID
	s.nextID++

	s.users = append(s.users, user)

	return nil
}

func (s *UserService) GetUsers() []model.User {
	return s.users
}

func (s *UserService) GetUsersByName(name string) (result []model.User) {
	for _, user := range s.users {
		if user.Name == name {
			result = append(result, user)
		}
	}
	return
}

func (s *UserService) GetUserByID(id int) (model.User, error) {
	for _, user := range s.users {
		if user.ID == id {
			return user, nil
		}
	}
	return model.User{}, ErrUserNotFound
}
