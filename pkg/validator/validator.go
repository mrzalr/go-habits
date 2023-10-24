package validator

import (
	"github.com/go-playground/validator/v10"
)

var (
	v *validator.Validate
)

func Load() {
	v = validator.New()
}

func ValidateStruct(target interface{}) error {
	err := v.Struct(target)
	if err == nil {
		return nil
	}

	ErrStructValidation = &validationError{
		Target: target,
		Errors: err.(validator.ValidationErrors),
	}

	return ErrStructValidation
}
