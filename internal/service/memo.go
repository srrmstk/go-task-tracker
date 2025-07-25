package service

import (
	"context"
	"go-task-tracker/internal/model"
	"go-task-tracker/internal/repository"
	"time"

	"github.com/google/uuid"
)

type MemoService struct {
	repo repository.MemoRepository
}

func NewMemoService(r repository.MemoRepository) *MemoService {
	return &MemoService{repo: r}
}

func (s *MemoService) GetAll(ctx context.Context) ([]model.Memo, error) {
	res, err := s.repo.GetAll(ctx)
	if res == nil && err == nil {
		return []model.Memo{}, nil
	}
	return res, err
}

func (s *MemoService) GetByID(ctx context.Context, id string) (model.Memo, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *MemoService) Create(ctx context.Context, dto model.MemoCreateDTO) (*model.Memo, error) {
	now := time.Now()

	m := &model.Memo{
		ID:          uuid.New(),
		Title:       dto.Title,
		Description: dto.Description,
		Score:       dto.Score,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := s.repo.Create(ctx, m); err != nil {
		return nil, err
	}

	return m, nil

}

func (s *MemoService) Update(ctx context.Context, id string, dto *model.MemoUpdateDTO) error {
	memo, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if dto.Title != nil {
		memo.Title = *dto.Title
	}
	if dto.Description != nil {
		memo.Description = *dto.Description
	}
	if dto.Score != nil {
		memo.Score = *dto.Score
	}

	memo.UpdatedAt = time.Now().UTC()

	return s.repo.Update(ctx, &memo)
}

func (s *MemoService) Delete(ctx context.Context, id string) error {
	return s.repo.Delete(ctx, id)
}
