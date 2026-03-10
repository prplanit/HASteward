package model

import "time"

// Envelope wraps every command result in a consistent schema for machine-readable output.
type Envelope[T any] struct {
	SchemaVersion string            `json:"schemaVersion"`
	Command       string            `json:"command"`
	RunID         string            `json:"runId,omitempty"`
	Success       bool              `json:"success"`
	Partial       bool              `json:"partial,omitempty"`
	ExitCode      int               `json:"exitCode"`
	Timestamp     time.Time         `json:"timestamp"`
	DurationMs    int64             `json:"durationMs,omitempty"`
	Warnings      []Warning         `json:"warnings,omitempty"`
	Errors        []StructuredError `json:"errors,omitempty"`
	Data          T                 `json:"data,omitempty"`
}

// NewEnvelope creates a new Envelope with standard fields populated.
func NewEnvelope[T any](command, runID string, data T) Envelope[T] {
	return Envelope[T]{
		SchemaVersion: "v1alpha1",
		Command:       command,
		RunID:         runID,
		Success:       true,
		ExitCode:      0,
		Timestamp:     time.Now().UTC(),
		Data:          data,
	}
}

// WithDuration sets the duration in milliseconds from a time.Duration.
func (e *Envelope[T]) WithDuration(d time.Duration) {
	e.DurationMs = d.Milliseconds()
}

// WithWarnings adds warnings to the envelope.
func (e *Envelope[T]) WithWarnings(warnings ...Warning) {
	e.Warnings = append(e.Warnings, warnings...)
}

// WithErrors marks the envelope as failed with structured errors.
func (e *Envelope[T]) WithErrors(exitCode int, errors ...StructuredError) {
	e.Success = false
	e.ExitCode = exitCode
	e.Errors = append(e.Errors, errors...)
}

// MarkPartial marks the envelope as a degraded/partial success (exit code 5).
func (e *Envelope[T]) MarkPartial() {
	e.Partial = true
	e.ExitCode = 5
}
