package apperror

import (
	"errors"
	"net/http"
)

type NewNotFoundErrorInput struct {
	Type    ErrorType
	Message string
	Err     error
}

func IsNotFoundError(err error) bool {
	if err == nil {
		return false
	}
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Type == NotFoundErrorType
}

func NewNotFoundError(input NewNotFoundErrorInput) *AppError {
	if input.Type == "" {
		input.Type = NotFoundErrorType
	}
	return &AppError{
		Type:    input.Type,
		Code:    http.StatusNotFound,
		Message: input.Message,
		Err:     input.Err,
	}
}
