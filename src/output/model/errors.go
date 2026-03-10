package model

import "fmt"

// Error categories for StructuredError.
const (
	CategoryUser         = "user"
	CategoryValidation   = "validation"
	CategoryEnvironment  = "environment"
	CategoryDependency   = "dependency"
	CategoryTransient    = "transient"
	CategoryPermission   = "permission"
	CategoryInternal     = "internal"
)

// Exit codes.
const (
	ExitSuccess          = 0
	ExitGenericFailure   = 1
	ExitInvalidUsage     = 2
	ExitValidationFailed = 3
	ExitDependencyFailed = 4
	ExitPartialSuccess   = 5
)

// StructuredError represents a machine-parseable error.
type StructuredError struct {
	Code      string         `json:"code"`
	Category  string         `json:"category"`
	Message   string         `json:"message"`
	Retryable bool           `json:"retryable"`
	Details   map[string]any `json:"details,omitempty"`
}

// Error implements the error interface.
func (e StructuredError) Error() string {
	return fmt.Sprintf("[%s] %s: %s", e.Category, e.Code, e.Message)
}

// NewError creates a StructuredError.
func NewError(code, category, message string) StructuredError {
	return StructuredError{
		Code:     code,
		Category: category,
		Message:  message,
	}
}

// NewRetryableError creates a retryable StructuredError.
func NewRetryableError(code, category, message string) StructuredError {
	return StructuredError{
		Code:      code,
		Category:  category,
		Message:   message,
		Retryable: true,
	}
}

// WithDetails returns a copy of the error with additional details.
func (e StructuredError) WithDetails(details map[string]any) StructuredError {
	e.Details = details
	return e
}
