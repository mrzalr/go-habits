package welcome

import "github.com/mrzalr/go-habits/internal/welcome/model"

type Usecase interface {
	GetMessage(name string) (model.Welcome, error)
}

type Repository interface {
	GetMessage(name string) (model.Welcome, error)
}
