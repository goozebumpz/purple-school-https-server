package auth

import (
	"errors"
	"purple-school/internal/user"
)

type Service struct {
	userRepository *user.Repository
}

func NewService(userRepo *user.Repository) *Service {
	return &Service{
		userRepository: userRepo,
	}
}

func (s *Service) Register(email, name, password string) (*user.User, error) {
	existedUser, err := s.userRepository.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	if existedUser != nil {
		return nil, errors.New(ErrUserAlreadyExist)
	}

	user := &user.User{
		Email:    email,
		Name:     name,
		Password: "",
	}

	user, err = s.userRepository.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}
