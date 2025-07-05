package repository

import (
	"testing"

	"github.com/daffarmd/task-manager/internal/model"
	_ "github.com/lib/pq"
)

func TestUserRepository_CreateAndFindByEmail(t *testing.T) {
	// Arrange
	email := "unittest1@example.com"
	user := model.User{
		Name:     "Test User",
		Email:    email,
		Password: "testpass",
	}

	// Act
	err := userRepoTest.Create(user)
	if err != nil {
		t.Fatalf("error creating user: %v", err)
	}

	// Assert
	got, err := userRepoTest.FindByEmail(email)
	if err != nil {
		t.Fatalf("error finding user: %v", err)
	}

	if got.Email != email {
		t.Errorf("expected email %s, got %s", email, got.Email)
	}
}
