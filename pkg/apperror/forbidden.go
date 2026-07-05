package apperror

import "errors"

func IsForbiddenError(err error) bool {
	if err == nil {
		return false
	}
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Type == ForbiddenErrorType
}

func NewForbiddenError(message string, err error) *AppError {
	return &AppError{
		Type:    ForbiddenErrorType,
		Code:    403,
		Message: message,
		Err:     err,
	}
}
