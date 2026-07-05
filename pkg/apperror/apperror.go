// Package apperror is the shared error model used by every Tentaculum service.
// Handlers translate any *AppError into a stable JSON envelope so callers see
// the same shape (type/code/message/data) regardless of which service replied.
package apperror

type AppError struct {
	Type    ErrorType `json:"type,omitempty"`
	Ref     string    `json:"code,omitempty"` // stable public error code, e.g. "AUTH-1001"
	Code    int       `json:"-"`              // HTTP status (internal)
	Message string    `json:"message,omitempty"`
	Err     error     `json:"-"`
	Data    any       `json:"data,omitempty"`
}

func (e *AppError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	if e.Err != nil {
		return e.Err.Error()
	}
	return string(e.Type)
}

func (e *AppError) Unwrap() error {
	if IsInternalServerError(e) {
		return e.Err
	}
	return nil
}

func (e *AppError) WithData(data any) *AppError {
	e.Data = data
	return e
}
