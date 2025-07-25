package service

import (
	"context"
	"go-task-tracker/internal/model"
	"go-task-tracker/internal/repository"
)

type CategoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(repo repository.CategoryRepository) *CategoryService {
	return &CategoryService{repo: repo}
}

func (s *CategoryService) GetAll(ctx context.Context) ([]model.Category, error) {
	res, err := s.repo.GetAll(ctx)
	if res == nil && err == nil {
		return []model.Category{}, nil
	}
	return res, err
}

func (s *CategoryService) GetByID(ctx context.Context, id int64) (model.Category, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *CategoryService) Create(ctx context.Context, c *model.Category) error {
	return s.repo.Create(ctx, c)
}

func (s *CategoryService) Update(ctx context.Context, id int64, c *model.CategoryUpdate) error {
	return s.repo.Update(ctx, id, c)
}

func (s *CategoryService) Delete(ctx context.Context, id int64) error {
	return s.repo.Delete(ctx, id)
}
