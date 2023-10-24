package usecase

import (
	"github.com/mrzalr/go-habits/internal/habit"
)

type usecase struct {
	repository habit.Repository
}

func New(repository habit.Repository) habit.Usecase {
	return &usecase{
		repository: repository,
	}
}
