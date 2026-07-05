package apperror

type ErrorType string

const (
	NotFoundErrorType       ErrorType = "not_found"
	InternalServerErrorType ErrorType = "internal_server_error"
	ValidationErrorType     ErrorType = "validation_error"
	ConflictErrorType       ErrorType = "conflict"
	ForbiddenErrorType      ErrorType = "forbidden"
	UnauthorizedErrorType   ErrorType = "unauthorized"
	RateLimitedErrorType    ErrorType = "rate_limited"
)
