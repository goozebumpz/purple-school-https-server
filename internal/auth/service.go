package auth

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
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

	if existedUser != nil {
		return nil, errors.New(ErrUserAlreadyExist)
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := &user.User{
		Email:    email,
		Name:     name,
		Password: string(hashedPassword),
	}

	user, err = s.userRepository.CreateUser(user)

	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *Service) Login(email, password string) (*user.User, error) {
	user, err := s.userRepository.FindByEmail(email)

	if err != nil {
		return nil, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return nil, err
	}

	return user, nil
}
