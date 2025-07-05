// internal/repository/repository_test.go
package repository

import (
	"database/sql"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var (
	testDB       *sql.DB
	userRepoTest UserRepository
	taskRepoTest TaskRepository
)

func TestMain(m *testing.M) {
	db, err := sql.Open("postgres", "host=localhost port=5432 user=devaja password=1q2w3e4r dbname=taskmanager sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}
	testDB = db

	userRepoTest = NewUserRepository(testDB)
	taskRepoTest = NewTaskRepository(testDB)

	code := m.Run()
	os.Exit(code)
}
