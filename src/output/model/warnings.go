package model

// Warning represents a typed, machine-parseable warning.
type Warning struct {
	Code    string         `json:"code"`
	Message string         `json:"message"`
	Details map[string]any `json:"details,omitempty"`
}

// NewWarning creates a Warning.
func NewWarning(code, message string) Warning {
	return Warning{
		Code:    code,
		Message: message,
	}
}

// WithDetails returns a copy of the warning with additional details.
func (w Warning) WithDetails(details map[string]any) Warning {
	w.Details = details
	return w
}
