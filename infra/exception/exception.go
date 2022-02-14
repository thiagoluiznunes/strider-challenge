package exception

import (
	"fmt"
	"net/http"
)

// GeneralError represents an input validation error
type GeneralError struct {
	Field   string `json:"field_name,omitempty"`
	Message string `json:"message,omitempty"`
	Code    int    `json:"code,omitempty"`
}

func (e *GeneralError) Error() string {
	if e.Field != "" {
		return fmt.Sprintf("%v: %v", e.Field, e.Message)
	}
	return fmt.Sprintf("%v", e.Message)
}

// NewValidationError returns a GeneralError instance with the provided parameters
func NewValidationError(field string, message string) *GeneralError {
	return &GeneralError{
		Field:   field,
		Message: message,
		Code:    http.StatusUnprocessableEntity,
	}
}

// NewNotFoundError returns a GeneralError instance with the provided parameters
func NewNotFoundError(err error) *GeneralError {
	return &GeneralError{
		Message: err.Error(),
		Code:    http.StatusNotFound,
	}
}

// NewApplicationError returns a GeneralError instance with the provided parameters
func NewApplicationError(err error) *GeneralError {
	return &GeneralError{
		Message: err.Error(),
		Code:    http.StatusServiceUnavailable,
	}
}

// NewConflictError returns a GeneralError instance with the provided parameters
func NewConflictError(message string) *GeneralError {
	return &GeneralError{
		Message: message,
		Code:    http.StatusConflict,
	}
}

// NewForbiddenError returns a GeneralError instance with the provided parameters
func NewForbiddenError(message string) *GeneralError {
	return &GeneralError{
		Message: message,
		Code:    http.StatusForbidden,
	}
}
