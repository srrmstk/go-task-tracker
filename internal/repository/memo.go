package repository

import (
	"context"
	"database/sql"
	"go-task-tracker/internal/model"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

type MemoRepository interface {
	GetAll(ctx context.Context) ([]model.Memo, error)
	GetByID(ctx context.Context, id int64) (model.Memo, error)
	Create(ctx context.Context, memo *model.Memo) error
	Update(ctx context.Context, id int64, memo *model.MemoUpdate) error
	Delete(ctx context.Context, id int64) error
}

func NewMemoRepository(db *sqlx.DB) MemoRepository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context) ([]model.Memo, error) {
	var memos []model.Memo
	query := "SELECT * FROM memos ORDER BY id"
	err := r.db.SelectContext(ctx, &memos, query)

	return memos, err
}

func (r *repository) GetByID(ctx context.Context, id int64) (model.Memo, error) {
	var memo model.Memo
	query := "SELECT * FROM memos WHERE id = $1"
	err := r.db.GetContext(ctx, &memo, query, id)
	return memo, err
}

func (r *repository) Create(ctx context.Context, memo *model.Memo) error {
	query := "INSERT INTO memos (title, description, score) VALUES ($1, $2, $3) returning id, created_at, updated_at"

	return r.db.
		QueryRowContext(ctx, query, memo.Title, memo.Description, memo.Score).
		Scan(&memo.ID, &memo.CreatedAt, &memo.UpdatedAt)
}

func (r *repository) Update(ctx context.Context, id int64, memo *model.MemoUpdate) error {
	query := `
			UPDATE memos 
			SET 
				title = COALESCE($2, title), 
				description = COALESCE($3, description), 
				score = COALESCE($4, score),
				updated_at = NOW() 
			WHERE id = $1`
	res, err := r.db.ExecContext(ctx, query, id, memo.Title, memo.Description, memo.Score)
	if err != nil {
		return err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int64) error {
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
