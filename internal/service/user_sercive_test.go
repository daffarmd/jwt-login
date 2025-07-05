package service

import (
	"errors"
	"testing"

	"github.com/daffarmd/task-manager/internal/model"
	"github.com/stretchr/testify/assert"
)

type mockUserRepo struct {
	users map[string]model.User
}

func (m *mockUserRepo) Create(user model.User) error {
	if _, exists := m.users[user.Email]; exists {
		return errors.New("user exists")
	}
	m.users[user.Email] = user
	return nil
}

func (m *mockUserRepo) FindByEmail(email string) (model.User, error) {
	user, exists := m.users[email]
	if !exists {
		return model.User{}, errors.New("not found")
	}
	return user, nil
}

func TestRegisterAndLogin(t *testing.T) {
	mock := &mockUserRepo{users: make(map[string]model.User)}
	service := NewUserService(mock)

	// Register user
	user := model.User{Name: "Daffa", Email: "daffa@example.com", Password: "secret123"}
	err := service.Register(user)
	assert.Nil(t, err)

	// Login success
	loggedIn, err := service.Login("daffa@example.com", "secret123")
	assert.Nil(t, err)
	assert.Equal(t, user.Email, loggedIn.Email)

	// Login failed
	_, err = service.Login("daffa@example.com", "wrongpass")
	assert.NotNil(t, err)
}
