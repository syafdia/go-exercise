package task

import "github.com/syafdia/clean-arch-ddd-cqrs-es/internal/domain"

const (
	ErrCodeInvalidTitleLength       domain.ErrorCode = "TASK_INVALID_TITLE_LENGTH"
	ErrCodeInvalidDescriptionLength domain.ErrorCode = "TASK_INVALID_DESCRIPTION"
	ErrCodeInvalidStateTransition   domain.ErrorCode = "TASK_INVALID_STATE_TRANSITION"
)
