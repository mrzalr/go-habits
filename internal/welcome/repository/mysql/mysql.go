package mysql

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	"github.com/mrzalr/go-habits/internal/welcome"
	"github.com/mrzalr/go-habits/internal/welcome/model"
)

type repository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) welcome.Repository {
	return &repository{
		db: db,
	}
}

func (r *repository) GetMessage(name string) (model.Welcome, error) {
	return model.Welcome{
		Message: fmt.Sprintf("welcome, %s !!!", name),
	}, nil
}
