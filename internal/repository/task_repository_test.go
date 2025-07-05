package repository

import (
	"testing"

	"github.com/daffarmd/task-manager/internal/model"
)

func TestTaskRepository_CreateAndGet(t *testing.T) {
	// Arrange
	task := model.Task{
		Title:  "Test Task",
		Done:   false,
		UserID: 1, // user id harus ada dulu
	}

	// Act
	err := taskRepoTest.Create(task)
	if err != nil {
		t.Fatalf("error creating task: %v", err)
	}

	// Act: Get all tasks
	tasks, err := taskRepoTest.GetAllByUserID(1)
	if err != nil {
		t.Fatalf("error getting tasks: %v", err)
	}

	// Assert
	found := false
	for _, t := range tasks {
		if t.Title == "Test Task" {
			found = true
			break
		}
	}

	if !found {
		t.Errorf("task not found in GetAllByUserID result")
	}
}
