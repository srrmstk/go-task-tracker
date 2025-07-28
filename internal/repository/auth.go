package repository

import (
	"database/sql"
	"go-task-tracker/internal/model"
	"time"

	"github.com/google/uuid"
	"github.com/jmoiron/sqlx"
	"github.com/redis/go-redis/v9"
	"golang.org/x/net/context"
)

type authRepository struct {
	db  *sqlx.DB
	rdb *redis.Client
}

type AuthRepository interface {
	Register(ctx context.Context, user *model.User) error
	GetUserByEmail(ctx context.Context, email string) (model.User, error)
	SetCode(ctx context.Context, id uuid.UUID, code string) error
	GetCode(ctx context.Context, id uuid.UUID) (string, error)
	Verify(ctx context.Context, id uuid.UUID, updatedAt time.Time) error
}

func NewAuthRepository(db *sqlx.DB, rdb *redis.Client) AuthRepository {
	return &authRepository{db: db, rdb: rdb}
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

func (r *authRepository) SetCode(ctx context.Context, id uuid.UUID, code string) error {
	ttl := time.Hour * 24
	if err := r.rdb.Set(ctx, id.String(), code, ttl).Err(); err != nil {
		return err
	}

	return nil
}

func (r *authRepository) GetCode(ctx context.Context, id uuid.UUID) (string, error) {
	res, err := r.rdb.Get(ctx, id.String()).Result()
	if err != nil {
		return "", err
	}

	return res, nil
}

func (r *authRepository) Verify(ctx context.Context, id uuid.UUID, updatedAt time.Time) error {
	query := `UPDATE users
			SET 
				confirmed = True, 
				updated_at = $2
			WHERE id = $1`

	res, err := r.db.ExecContext(ctx, query, id, updatedAt)
	if err != nil {
		return err
	}

	rows, _ := res.RowsAffected()
	if rows == 0 {
		return sql.ErrNoRows
	}

	if err := r.rdb.Del(ctx, id.String()).Err(); err != nil {
		return err
	}

	return nil
}
