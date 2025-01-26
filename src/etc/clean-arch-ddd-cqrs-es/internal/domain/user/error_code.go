package user

import "github.com/syafdia/clean-arch-ddd-cqrs-es/internal/domain"

const (
	ErrCodeInvalidPasswordLength domain.ErrorCode = "USER-INVALID_PASSWORD_LENGTH"
	ErrCodeInvalidEmailFormat    domain.ErrorCode = "USER-INVALID_EMAIL_FORMAT"
)
