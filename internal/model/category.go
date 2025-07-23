package model

import "time"

type Category struct {
	ID        int64     `db:"id" json:"id"`
	Title     string    `db:"title" json:"title" validate:"required"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type CategoryUpdate struct {
	Title string `db:"title" json:"title" validate:"required"`
}
