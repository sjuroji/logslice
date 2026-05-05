package config

import (
	"errors"
	"fmt"
)

// ValidationError represents a configuration validation error.
type ValidationError struct {
	Field   string
	Message string
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("config: invalid %s: %s", e.Field, e.Message)
}

// Validate checks the parsed Config for logical consistency and returns
// an error if any field combination is invalid.
func Validate(cfg *Config) error {
	if cfg == nil {
		return errors.New("config: nil config")
	}

	if cfg.MaxLines < 0 {
		return &ValidationError{
			Field:   "max-lines",
			Message: "must be non-negative",
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

	if cfg.End.IsZero() && !cfg.Start.IsZero() {
		// start without end is fine — open-ended range
	}

	if !cfg.End.IsZero() && cfg.Start.IsZero() {
		// end without start is fine — open-ended range
	}

	return nil
}
