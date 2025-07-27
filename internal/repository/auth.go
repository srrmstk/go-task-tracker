package repository

import (
	"database/sql"
	"go-task-tracker/internal/model"

	"github.com/jmoiron/sqlx"
	"golang.org/x/net/context"
)

type authRepository struct {
	db *sqlx.DB
}

type AuthRepository interface {
	Register(ctx context.Context, user *model.User) error
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
}

func NewAuthRepository(db *sqlx.DB) AuthRepository {
	return &authRepository{db: db}
}

func (r *authRepository) Register(ctx context.Context, user *model.User) error {
	query := `INSERT INTO users (id, username, password, email, created_at, updated_at)
				VALUES ($1, $2, $3, $4, $5, $6)`

	res, err := r.db.ExecContext(ctx, query, user.ID, user.Username, user.Password, user.Email, user.CreatedAt, user.UpdatedAt)
	if err != nil {
		return err
	}

	if rows, _ := res.RowsAffected(); rows == 0 {
		return sql.ErrNoRows
	}

	return nil
}

func (r *authRepository) GetUserByEmail(ctx context.Context, email string) (model.User, error) {
	var user model.User
	query := "SELECT * FROM users WHERE email = $1"
	err := r.db.GetContext(ctx, &user, query, email)
	return user, err
}
