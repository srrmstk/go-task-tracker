package repository

import (
	"context"
	"database/sql"
	"go-task-tracker/internal/model"

	"github.com/jmoiron/sqlx"
)

type categoryRepository struct {
	db *sqlx.DB
}

type CategoryRepository interface {
	GetAll(ctx context.Context) ([]model.Category, error)
	GetByID(ctx context.Context, id int64) (model.Category, error)
	Create(ctx context.Context, category *model.Category) error
	Update(ctx context.Context, id int64, category *model.CategoryUpdate) error
	Delete(ctx context.Context, id int64) error
}

func NewCategoryRepository(db *sqlx.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetAll(ctx context.Context) ([]model.Category, error) {
	var res []model.Category
	query := "SELECT * FROM categories"
	err := r.db.SelectContext(ctx, res, query)
	return res, err
}

func (r *categoryRepository) GetByID(ctx context.Context, id int64) (model.Category, error) {
	var res model.Category
	query := "SELECT * FROM categories WHERE id = $1"
	err := r.db.SelectContext(ctx, res, query, id)
	return res, err
}

func (r *categoryRepository) Create(ctx context.Context, category *model.Category) error {
	query := "INSERT INTO categories (title) VALUES ($1) returning id, created_at, updated_at"
	return r.db.QueryRowContext(ctx, query, category.Title).Scan(&category.ID, &category.CreatedAt, &category.UpdatedAt)
}

func (r *categoryRepository) Update(ctx context.Context, id int64, category *model.CategoryUpdate) error {
	query := `
		UPDATE categories 
		SET
			title = COALESCE($2, title),
			updated_at = NOW()
		WHERE id = $1`
	res, err := r.db.ExecContext(ctx, query, id, category.Title)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil

}

func (r *categoryRepository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM categories WHERE id = $1"
	res, err := r.db.ExecContext(ctx, query)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
