package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var ErrStructValidation error

type validationError struct {
	Target interface{}
	Errors validator.ValidationErrors
}

func (e *validationError) Error() string {
	msgs := []string{}
	for _, vErrs := range e.Errors {
		msgs = append(msgs, sugar(e.Target, vErrs))
	}

	return strings.Join(msgs, "; ")
}

func sugar(target interface{}, err validator.FieldError) string {
	jsonTagField := err.Field()

	field, ok := reflect.TypeOf(target).FieldByName(err.Field())
	if ok {
		jsonTagField = field.Tag.Get("json")
	}

	switch err.Tag() {
	case "required":
		return fmt.Sprintf("field %s cannot be empty", jsonTagField)
	case "min":
		return fmt.Sprintf("field %s must be have minimum %s character", jsonTagField, err.Param())
	default:
		return fmt.Sprintf("field %s validation failed on tag %s", jsonTagField, err.Tag())
	}
}
