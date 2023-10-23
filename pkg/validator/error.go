package validator

import (
	"fmt"
	"strings"

	"github.com/go-playground/validator/v10"
)

var ErrStructValidation error

type validationError struct {
	Errors validator.ValidationErrors
}

func (e *validationError) Error() string {
	msgs := []string{}
	for _, vErrs := range e.Errors {
		msgs = append(msgs, sugar(vErrs))
	}

	return strings.Join(msgs, "; ")
}

func sugar(err validator.FieldError) string {
	switch err.Tag() {
	case "required":
		return fmt.Sprintf("field %s cannot be empty", err.Field())
	case "min":
		return fmt.Sprintf("field %s must be have minimum %s character", err.Field(), err.Param())
	default:
		return fmt.Sprintf("field %s validation failed on tag %s", err.Field(), err.Tag())
	}
}
