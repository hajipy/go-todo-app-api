package handler

import (
	"context"

	"github.com/hajipy/go_todo_app/entity"
)

//go:generate go run github.com/matryer/moq -out moq_test.go . AddTaskService ListTasksService RegisterUserService
type AddTaskService interface {
	AddTask(ctx context.Context, title string) (*entity.Task, error)
}

type ListTasksService interface {
	ListTasks(ctx context.Context) (entity.Tasks, error)
}

type RegisterUserService interface {
	RegisterUser(ctx context.Context, name, password, role string) (*entity.User, error)
}
