package model

import (
	"time"

	"github.com/google/uuid"
)

type Memo struct {
	ID          uuid.UUID `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	Score       int       `db:"score" json:"score"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type MemoCreateDTO struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	Score       int    `json:"score" validate:"required,gte=1,lte=10"`
}

type MemoUpdateDTO struct {
	Title       *string `json:"title,omitempty" validate:"omitempty,min=1"`
	Description *string `json:"description,omitempty" validate:"omitempty"`
	Score       *int    `json:"score" validate:"omitempty,gte=1,lte=10"`
}
