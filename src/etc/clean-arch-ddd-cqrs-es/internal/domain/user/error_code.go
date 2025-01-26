package user

import "github.com/syafdia/clean-arch-ddd-cqrs-es/internal/domain"

const (
	ErrCodeInvalidPasswordLength domain.ErrorCode = "USER_INVALID_PASSWORD_LENGTH"
	ErrCodeInvalidEmailFormat    domain.ErrorCode = "USER_INVALID_EMAIL_FORMAT"
)
