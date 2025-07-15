package service

import (
	"context"
	"errors"
	"go-task-tracker/internal/model"
)

type TaskRepository interface {
	GetAll(ctx context.Context) ([]model.Task, error)
	GetByID(ctx context.Context, id string) (model.Task, error)
	Create(ctx context.Context, task *model.Task) error
}

type TaskService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) *TaskService {
	return &TaskService{repo: r}
}

func (s *TaskService) GetAll(ctx context.Context) ([]model.Task, error) {
	return s.repo.GetAll(ctx)
}

func (s *TaskService) GetByID(ctx context.Context, id string) (model.Task, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *TaskService) Create(ctx context.Context, t *model.Task) error {
	if t.Title == "" {
		return errors.New("title is required")
	}
	return s.repo.Create(ctx, t)
}
