package handler

import (
	"context"

	"github.com/hajipy/go_todo_app/entity"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . AddTaskService ListTasksService
type AddTaskService interface {
	AddTask(ctx context.Context, title string) (*entity.Task, error)
}

type ListTasksService interface {
	ListTask(ctx context.Context) (entity.Tasks, error)
}
