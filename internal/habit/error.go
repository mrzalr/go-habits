package habit

import "fmt"

var (
	ErrDataNotFound   = fmt.Errorf("habit: data not found")
	ErrInternalServer = fmt.Errorf("habit: internal server error")
)
