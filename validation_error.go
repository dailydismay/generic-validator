package validator

import "errors"

type ValidationError interface {
	error
	Field() string
	Message() string
}

func IsValidationError(err error) bool {
	var ve ValidationError

	return errors.As(err, &ve)
}

func MustAsValidationError(err error) ValidationError {
	var ve ValidationError

	if !errors.As(err, &ve) {
		panic("unexpected error passed as validation error")
	}

	return ve
}

type validationError struct {
	field   string
	message string
}

func NewValidationError(field, msg string) ValidationError {
	return &validationError{field, msg}
}

func (ve validationError) Error() string {
	return ve.message
}

func (ve validationError) Field() string {
	return ve.field
}

func (ve validationError) Message() string {
	return ve.message
}
