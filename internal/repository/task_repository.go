package repository

import (
	"database/sql"

	"github.com/daffarmd/task-manager/internal/model"
)

type TaskRepository interface {
	Create(task model.Task) error
	GetAllByUserID(userID int) ([]model.Task, error)
}

type taskRepo struct {
	DB *sql.DB
}

func NewTaskRepository(db *sql.DB) TaskRepository {
	return &taskRepo{DB: db}
}

func (r *taskRepo) Create(task model.Task) error {
	_, err := r.DB.Exec("INSERT INTO tasks (title, done, user_id) VALUES ($1, $2, $3)", task.Title, task.Done, task.UserID)
	return err
}

func (r *taskRepo) GetAllByUserID(userID int) ([]model.Task, error) {
	rows, err := r.DB.Query("SELECT id, title, done, user_id FROM tasks WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tasks []model.Task
	for rows.Next() {
		var t model.Task
		if err := rows.Scan(&t.ID, &t.Title, &t.Done, &t.UserID); err != nil {
			return nil, err
		}
		tasks = append(tasks, t)
	}
	return tasks, nil
}
