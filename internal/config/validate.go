package config

import (
	"errors"
	"fmt"
)

// ValidationError represents a configuration validation failure.
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("config: invalid field %q: %s", e.Field, e.Message)
}

// Validate checks the parsed Config for logical consistency and returns
// an error if any constraint is violated.
func Validate(cfg *Config) error {
	if cfg == nil {
		return errors.New("config: nil config")
	}

	if cfg.MaxLines < 0 {
		return &ValidationError{
			Field:   "max-lines",
			Message: "must be >= 0",
		}
	}

	if !cfg.Start.IsZero() && !cfg.End.IsZero() {
		if cfg.End.Before(cfg.Start) {
			return &ValidationError{
				Field:   "end",
				Message: "must not be before start",
			}
		}
	}

	if cfg.Workers < 1 {
		return &ValidationError{
			Field:   "workers",
			Message: "must be >= 1",
		}
	}

	return nil
}
