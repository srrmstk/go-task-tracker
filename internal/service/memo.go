package service

import (
	"context"
	"errors"
	"go-task-tracker/internal/model"
	"go-task-tracker/internal/repository"
)

type MemoService struct {
	repo repository.MemoRepository
}

func NewMemoService(r repository.MemoRepository) *MemoService {
	return &MemoService{repo: r}
}

func (s *MemoService) GetAll(ctx context.Context) ([]model.Memo, error) {
	return s.repo.GetAll(ctx)
}

func (s *MemoService) GetByID(ctx context.Context, id int64) (model.Memo, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *MemoService) Create(ctx context.Context, m *model.Memo) error {
	if m.Title == "" {
		return errors.New("title is required")
	}
	return s.repo.Create(ctx, m)
}

func (s *MemoService) Update(ctx context.Context, id int64, m *model.MemoUpdate) error {
	return s.repo.Update(ctx, id, m)
}

func (s *MemoService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
