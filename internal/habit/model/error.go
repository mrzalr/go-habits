package model

import (
	"errors"
)

var (
	ErrDataNotFound   = errors.New("habit: data not found")
	ErrAlreadyStarted = errors.New("habit: activity already started")
	ErrBadRequest     = errors.New("habit: bad request")
)
