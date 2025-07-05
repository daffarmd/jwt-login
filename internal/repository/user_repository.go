package repository

import (
	"database/sql"

	"github.com/daffarmd/task-manager/internal/model"
)

type UserRepository interface {
	Create(user model.User) error
	FindByEmail(email string) (model.User, error)
}

type userRepo struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepo{DB: db}
}

func (r *userRepo) Create(user model.User) error {
	_, err := r.DB.Exec("INSERT INTO users (name, email, password) VALUES ($1, $2, $3)", user.Name, user.Email, user.Password)
	return err
}

func (r *userRepo) FindByEmail(email string) (model.User, error) {
	var user model.User
	err := r.DB.QueryRow("SELECT id, name, email, password FROM users WHERE email = $1", email).
		Scan(&user.ID, &user.Name, &user.Email, &user.Password)
	return user, err
}
