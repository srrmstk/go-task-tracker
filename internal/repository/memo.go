package repository

import (
	"context"
	"database/sql"
	"go-task-tracker/internal/model"

	"github.com/jmoiron/sqlx"
)

type memoRepository struct {
	db *sqlx.DB
}

type MemoRepository interface {
	GetAll(ctx context.Context) ([]model.Memo, error)
	GetByID(ctx context.Context, id string) (model.Memo, error)
	Create(ctx context.Context, memo *model.Memo) error
	Update(ctx context.Context, memo *model.Memo) error
	Delete(ctx context.Context, id string) error
}

func NewMemoRepository(db *sqlx.DB) MemoRepository {
	return &memoRepository{db: db}
}

func (r *memoRepository) GetAll(ctx context.Context) ([]model.Memo, error) {
	var memos []model.Memo
	query := "SELECT * FROM memos ORDER BY id"
	err := r.db.SelectContext(ctx, &memos, query)
	return memos, err
}

func (r *memoRepository) GetByID(ctx context.Context, id string) (model.Memo, error) {
	var memo model.Memo
	query := "SELECT * FROM memos WHERE id = $1"
	err := r.db.GetContext(ctx, &memo, query, id)
	return memo, err
}

func (r *memoRepository) Create(ctx context.Context, memo *model.Memo) error {
	query := `INSERT INTO memos (id, title, description, score, created_at, updated_at) 
				VALUES ($1, $2, $3, $4, $5, $6) 
				RETURNING id, created_at, updated_at`
	return r.db.
		QueryRowContext(ctx, query, memo.ID, memo.Title, memo.Description, memo.Score, memo.CreatedAt, memo.UpdatedAt).
		Scan(&memo.ID, &memo.CreatedAt, &memo.UpdatedAt)
}

func (r *memoRepository) Update(ctx context.Context, memo *model.Memo) error {
	query := `
			UPDATE memos 
			SET 
				title = COALESCE($2, title), 
				description = COALESCE($3, description), 
				score = COALESCE($4, score),
				updated_at = $5
			WHERE id = $1`
	res, err := r.db.ExecContext(ctx, query, memo.ID, memo.Title, memo.Description, memo.Score, memo.UpdatedAt)
	if err != nil {
		return err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *memoRepository) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM memos WHERE id = $1"
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
