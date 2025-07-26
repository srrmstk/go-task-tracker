package service

import (
	"context"
	"go-task-tracker/internal/model"
	"go-task-tracker/internal/repository"
	"time"

	"github.com/google/uuid"
)

type MemoService struct {
	memoRepo repository.MemoRepository
	cateRepo repository.CategoryRepository
}

func NewMemoService(mr repository.MemoRepository, cr repository.CategoryRepository) *MemoService {
	return &MemoService{memoRepo: mr, cateRepo: cr}
}

func (s *MemoService) GetAll(ctx context.Context) ([]model.Memo, error) {
	res, err := s.memoRepo.GetAll(ctx)
	if res == nil && err == nil {
		return []model.Memo{}, nil
	}
	return res, err
}

func (s *MemoService) GetByID(ctx context.Context, id uuid.UUID) (model.Memo, error) {
	return s.memoRepo.GetByID(ctx, id)
}

func (s *MemoService) Create(ctx context.Context, dto model.MemoCreateDTO) (*model.Memo, error) {
	now := time.Now().UTC()

	_, err := s.cateRepo.GetByID(ctx, dto.CategoryID)
	if err != nil {
		return nil, err
	}

	m := &model.Memo{
		ID:          uuid.New(),
		Title:       dto.Title,
		Description: dto.Description,
		Score:       dto.Score,
		CategoryID:  dto.CategoryID,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	if err := s.memoRepo.Create(ctx, m); err != nil {
		return nil, err
	}

	return m, nil

}

func (s *MemoService) Update(ctx context.Context, id uuid.UUID, dto *model.MemoUpdateDTO) error {
	memo, err := s.memoRepo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	if dto.CategoryID != nil {
		_, err := s.cateRepo.GetByID(ctx, *dto.CategoryID)
		if err != nil {
			return err
		}
		memo.CategoryID = *dto.CategoryID
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

	return s.memoRepo.Update(ctx, &memo)
}

func (s *MemoService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.memoRepo.Delete(ctx, id)
}
