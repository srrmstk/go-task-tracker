package service

import (
	"context"
	"go-task-tracker/internal/model"
	"go-task-tracker/internal/repository"
	"time"

	"github.com/google/uuid"
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

func (s *CategoryService) GetByID(ctx context.Context, id uuid.UUID) (model.Category, error) {
	return s.repo.GetByID(ctx, id)
}

func (s *CategoryService) Create(ctx context.Context, dto model.CategoryCreateDTO) (*model.Category, error) {
	now := time.Now().UTC()

	model := &model.Category{
		ID:        uuid.New(),
		Title:     dto.Title,
		CreatedAt: now,
		UpdatedAt: now,
	}

	err := s.repo.Create(ctx, model)
	if err != nil {
		return nil, err
	}

	return model, nil
}

func (s *CategoryService) Update(ctx context.Context, id uuid.UUID, c *model.CategoryUpdateDTO) error {
	category, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return err
	}

	category.Title = c.Title
	category.UpdatedAt = time.Now().UTC()

	return s.repo.Update(ctx, &category)
}

func (s *CategoryService) Delete(ctx context.Context, id uuid.UUID) error {
	return s.repo.Delete(ctx, id)
}
