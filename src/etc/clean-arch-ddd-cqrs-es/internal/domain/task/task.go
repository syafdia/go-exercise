package task

import (
	"time"

	"github.com/syafdia/clean-arch-ddd-cqrs-es/internal/domain"
	"github.com/syafdia/clean-arch-ddd-cqrs-es/internal/domain/user"
)

type TaskID int64

type TaskStatus string

const (
	TaskStatusNotStarted TaskStatus = "NOT_STARTED"
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	TaskStatusInReview   TaskStatus = "IN_REVIEW"
	TaskStatusCompleted  TaskStatus = "COMPLETED"
)

type Task struct {
	ID          TaskID
	AuthorID    user.UserID
	AssigneeID  user.UserID
	Title       Title
	Description Description
	Status      TaskStatus
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

func NewTask(
	authorID user.UserID,
	title Title,
	description Description,
	createdAt time.Time,
) (*Task, error) {
	err := title.Validate()
	if err != nil {
		return nil, err
	}

	err = description.Validate()
	if err != nil {
		return nil, err
	}

	return &Task{
		Title:       title,
		Description: description,
		CreatedAt:   createdAt,
	}, nil

}

type Title string

func (t Title) Validate() error {
	if len(t) < 4 {
		return domain.NewValidationError(ErrCodeInvalidTitleLength, "must be at least 6 characters")
	}

	if len(t) > 32 {
		return domain.NewValidationError(ErrCodeInvalidTitleLength, "the maximum value is 32 characters")
	}

	return nil
}

type Description string

func (d Description) Validate() error {
	if len(d) > 256 {
		return domain.NewValidationError(ErrCodeInvalidDescriptionLength, "the maximum value is 256 characters")
	}

	return nil
}
