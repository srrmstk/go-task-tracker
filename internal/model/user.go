package model

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Username  string    `db:"username" json:"username"`
	Password  string    `db:"password" json:"password"`
	Email     string    `db:"email" json:"email"`
	Confirmed bool      `db:"confirmed" json:"confirmed"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type UserLoginDTO struct {
	Email    string `json:"username" validate:"required,email"`
	Password string `json:"password" validate:"required,gte=8"`
}

type UserRegisterDTO struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required,gte=8"`
	Email    string `json:"email" validate:"required,email"`
}
