package repository

import (
	"context"
	"go-task-tracker/internal/model"
	"go-task-tracker/internal/service"

	"github.com/jmoiron/sqlx"
)

type repository struct {
	db *sqlx.DB
}

func NewTaskRepository(db *sqlx.DB) service.TaskRepository {
	return &repository{db: db}
}

func (r *repository) GetAll(ctx context.Context) ([]model.Task, error) {
	var tasks []model.Task
	query := "SELECT * FROM tasks ORDER BY id"
	err := r.db.SelectContext(ctx, &tasks, query)

	return tasks, err
}

func (r *repository) GetByID(ctx context.Context, id string) (model.Task, error) {
	var task model.Task
	query := "SELECT * FROM tasks Where id = $1"
	err := r.db.SelectContext(ctx, &task, query, id)
	return task, err
}

func (r *repository) Create(ctx context.Context, task *model.Task) error {
	query := "INSERT INTO tasks (title, description) VALUES ($1, $2) returning id, created_at, updated_at"

	return r.db.
		QueryRowContext(ctx, query, task.Title, task.Description).
		Scan(&task.ID, &task.CreatedAt, &task.UpdatedAt)
}
