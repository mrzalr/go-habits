package habit

import (
	"errors"
)

var (
	ErrDataNotFound   = errors.New("habit: data not found")
	ErrInternalServer = errors.New("habit: internal server error")
	ErrAlreadyStarted = errors.New("habit: activity already started")
)
