package model

import "time"

type Memo struct {
	ID          int64     `db:"id" json:"id"`
	Title       string    `db:"title" json:"title" validate:"required"`
	Description string    `db:"description" json:"description"`
	Score       int       `db:"score" json:"score" validate:"required,gte=1,lte=10"`
	CreatedAt   time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   time.Time `db:"updated_at" json:"updated_at"`
}

type MemoUpdate struct {
	Title       *string `json:"title,omitempty" validate:"omitempty,min=1"`
	Description *string `json:"description,omitempty" validate:"omitempty"`
	Score       *int    `json:"score" validate:"omitempty,gte=1,lte=10"`
}
