package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/daffarmd/task-manager/internal/model"
	_ "github.com/lib/pq"
)

var testRepo UserRepository

func TestMain(m *testing.M) {
	// Connect ke test DB
	db, err := sql.Open("postgres", "host=localhost port=5432 user=devaja password=1q2w3e4r dbname=taskmanager sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	testRepo = NewUserRepository(db)

	code := m.Run()
	os.Exit(code)
}

func TestUserRepository_CreateAndFindByEmail(t *testing.T) {
	// Arrange
	email := "unittest@example.com"
	user := model.User{
		Name:     "Test User",
		Email:    email,
		Password: "testpass",
	}

	// Act
	err := testRepo.Create(user)
	if err != nil {
		t.Fatalf("error creating user: %v", err)
	}

	// Assert
	got, err := testRepo.FindByEmail(email)
	if err != nil {
		t.Fatalf("error finding user: %v", err)
	}

	if got.Email != email {
		t.Errorf("expected email %s, got %s", email, got.Email)
	}
}
