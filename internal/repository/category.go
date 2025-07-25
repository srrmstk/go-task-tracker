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
	GetByID(ctx context.Context, id string) (model.Category, error)
	Create(ctx context.Context, category *model.Category) error
	Update(ctx context.Context, category *model.Category) error
	Delete(ctx context.Context, id string) error
}

func NewCategoryRepository(db *sqlx.DB) CategoryRepository {
	return &categoryRepository{db: db}
}

func (r *categoryRepository) GetAll(ctx context.Context) ([]model.Category, error) {
	var res []model.Category
	query := "SELECT * FROM categories"
	err := r.db.SelectContext(ctx, &res, query)
	return res, err
}

func (r *categoryRepository) GetByID(ctx context.Context, id string) (model.Category, error) {
	var res model.Category
	query := "SELECT * FROM categories WHERE id = $1"
	err := r.db.GetContext(ctx, &res, query, id)
	return res, err
}

func (r *categoryRepository) Create(ctx context.Context, category *model.Category) error {
	query := `INSERT INTO categories (id, title, created_at, updated_at) 
				VALUES ($1, $2, $3, $4) 
				RETURNING id, created_at, updated_at`
	return r.db.
		QueryRowContext(ctx, query, category.ID, category.Title, category.CreatedAt, category.UpdatedAt).
		Scan(&category.ID, &category.CreatedAt, &category.UpdatedAt)
}

func (r *categoryRepository) Update(ctx context.Context, category *model.Category) error {
	query := `
		UPDATE categories 
		SET
			title = COALESCE($2, title),
			updated_at = $3
		WHERE id = $1`
	res, err := r.db.ExecContext(ctx, query, category.ID, category.Title, category.UpdatedAt)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil

}

func (r *categoryRepository) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM categories WHERE id = $1"
	res, err := r.db.ExecContext(ctx, query, id)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}
