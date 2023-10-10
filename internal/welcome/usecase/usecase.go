package usecase

import (
	"github.com/mrzalr/go-habits/internal/welcome"
	"github.com/mrzalr/go-habits/internal/welcome/model"
)

type usecase struct {
	repository welcome.Repository
}

func New(repository welcome.Repository) welcome.Usecase {
	return &usecase{
		repository: repository,
	}
}

func (u *usecase) GetMessage(name string) (model.Welcome, error) {
	return u.repository.GetMessage(name)
}
