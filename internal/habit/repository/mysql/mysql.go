package mysql

import (
	"github.com/jmoiron/sqlx"
	"github.com/mrzalr/go-habits/internal/habit"
)

type queryParams map[string]any

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) habit.Repository {
	return &repository{
		db: db,
	}
}
