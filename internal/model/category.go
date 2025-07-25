package model

import (
	"time"

	"github.com/google/uuid"
)

type Category struct {
	ID        uuid.UUID `db:"id" json:"id"`
	Title     string    `db:"title" json:"title"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type CategoryCreateDTO struct {
	Title string `json:"title" validate:"required"`
}

type CategoryUpdateDTO struct {
	Title string `json:"title" validate:"required"`
}
