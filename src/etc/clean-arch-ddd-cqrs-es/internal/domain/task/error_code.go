package task

import "github.com/syafdia/clean-arch-ddd-cqrs-es/internal/domain"

const (
	ErrCodeInvalidTitleLength       domain.ErrorCode = "TASK-INVALID_TITLE_LENGTH"
	ErrCodeInvalidDescriptionLength domain.ErrorCode = "TASK-INVALID_DESCRIPTION"
	ErrCodeInvalidStateTransition   domain.ErrorCode = "TASK-INVALID_STATE_TRANSITION"
	ErrCodeInvalidTagNameLength     domain.ErrorCode = "TASK-INVALID_TAG_NAME_LENGTH"
	ErrCodeTagAlreadyExists         domain.ErrorCode = "TASK-TAG_ALREADY_EXISTS"
)
