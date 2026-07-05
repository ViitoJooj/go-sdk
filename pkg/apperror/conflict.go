package apperror

import "errors"

func IsConflictError(err error) bool {
	if err == nil {
		return false
	}
	var appErr *AppError
	return errors.As(err, &appErr) && appErr.Type == ConflictErrorType
}

func NewConflictError(message string, err error) *AppError {
	return &AppError{
		Type:    ConflictErrorType,
		Code:    409,
		Message: message,
		Err:     err,
	}
}
