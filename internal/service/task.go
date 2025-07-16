package service

import (
	"context"
	"errors"
	"go-task-tracker/internal/model"
	"go-task-tracker/internal/repository"
)

type TaskService struct {
	repo repository.TaskRepository
}

func NewTaskService(r repository.TaskRepository) *TaskService {
	return &TaskService{repo: r}
}

func (s *TaskService) GetAll(ctx context.Context) ([]model.Task, error) {
	return s.repo.GetAll(ctx)
}

func (s *TaskService) GetByID(ctx context.Context, id int64) (model.Task, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *TaskService) Create(ctx context.Context, t *model.Task) error {
	if t.Title == "" {
		return errors.New("title is required")
	}
	return s.repo.Create(ctx, t)
}

func (s *TaskService) Update(ctx context.Context, id int64, t *model.TaskUpdate) error {
	if t.Title == nil && t.Description == nil {
		return errors.New("nothing to update")
	}
	return s.repo.Update(ctx, id, t)
}
