package domain

import "github.com/syafdia/clean-arch-ddd-cqrs-es/pkg/errr"

type ErrorCode string

type ValidationError struct {
	errr.Err
}

func NewValidationError(code ErrorCode, message string) *ValidationError {
	return &ValidationError{*errr.NewErr(string(code), message)}
}
