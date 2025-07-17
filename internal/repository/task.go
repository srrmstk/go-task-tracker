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

type TaskRepository interface {
	GetAll(ctx context.Context) ([]model.Task, error)
	GetByID(ctx context.Context, id int64) (model.Task, error)
	Create(ctx context.Context, task *model.Task) error
	Update(ctx context.Context, id int64, task *model.TaskUpdate) error
	Delete(ctx context.Context, id int64) error
}

func NewTaskRepository(db *sqlx.DB) TaskRepository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context) ([]model.Task, error) {
	var tasks []model.Task
	query := "SELECT * FROM tasks ORDER BY id"
	err := r.db.SelectContext(ctx, &tasks, query)

	return tasks, err
}

func (r *repository) GetByID(ctx context.Context, id int64) (model.Task, error) {
	var task model.Task
	query := "SELECT * FROM tasks WHERE id = $1"
	err := r.db.GetContext(ctx, &task, query, id)
	return task, err
}

func (r *repository) Create(ctx context.Context, task *model.Task) error {
	query := "INSERT INTO tasks (title, description) VALUES ($1, $2) returning id, created_at, updated_at"

	return r.db.
		QueryRowContext(ctx, query, task.Title, task.Description).
		Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
}

func (r *repository) Update(ctx context.Context, id int64, task *model.TaskUpdate) error {
	query := `
			UPDATE tasks 
			SET 
				title = COALESCE($2, title), 
				description = COALESCE($3, description), 
				updated_at = NOW() 
			WHERE id = $1`
	res, err := r.db.ExecContext(ctx, query, id, task.Title, task.Description)
	if err != nil {
		return err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *repository) Delete(ctx context.Context, id int64) error {
	query := "DELETE FROM tasks WHERE id = $1"
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
