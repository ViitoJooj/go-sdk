package apperror

import (
	"errors"
	"net/http"
)

type InternalServerErrorInput struct {
	Type ErrorType
	Err  error
}

func IsInternalServerError(err error) bool {
	if err == nil {
		return false
	}
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Type == InternalServerErrorType
}

func NewInternalServerError(input InternalServerErrorInput) *AppError {
	if input.Type == "" {
		input.Type = InternalServerErrorType
	}

	return &AppError{
		Type:    input.Type,
		Code:    http.StatusInternalServerError,
		Message: "internal server error",
		Err:     input.Err,
	}
}
