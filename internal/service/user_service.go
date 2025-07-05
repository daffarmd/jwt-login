package service

import (
	"errors"

	"github.com/daffarmd/task-manager/internal/model"
	"github.com/daffarmd/task-manager/internal/repository"
	"golang.org/x/crypto/bcrypt"
)

type UserService interface {
	Register(user model.User) error
	Login(email, password string) (model.User, error)
}

type userService struct {
	repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(user model.User) error {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	user.Password = string(hash)

	// save to DB
	return s.repo.Create(user)
}

func (s *userService) Login(email, password string) (model.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return model.User{}, err
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return model.User{}, errors.New("Invalid credentials")
	}

	return user, nil
}
