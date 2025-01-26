package task

import (
	"fmt"
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
	Tags        []Tag
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

func (t *Task) ChangeStatus(newStatus TaskStatus) error {
	// TODO
	return nil
}

func (t *Task) AddTag(newTag Tag, updatedAt time.Time) error {
	for _, tag := range t.Tags {
		if tag.ID == newTag.ID {
			return domain.NewValidationError(
				ErrCodeTagAlreadyExists,
				fmt.Sprintf("tag %s already exists on this task", newTag.Name))
		}
	}

	t.Tags = append(t.Tags, newTag)
	t.UpdatedAt = updatedAt

	return nil
}

func (t *Task) RemoveTag(newTag Tag) error {
	// TODO
	return domain.ErrNotImplemented
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

type TagID int64

type Tag struct {
	ID        TagID
	Name      TagName
	CreatedAt time.Time
}

func NewTag(name TagName, createdAt time.Time) (*Tag, error) {
	err := name.Validate()
	if err != nil {
		return nil, err
	}

	return &Tag{
		Name:      name,
		CreatedAt: createdAt,
	}, nil
}

type TagName string

func (tn TagName) Validate() error {
	if len(tn) < 2 {
		return domain.NewValidationError(ErrCodeInvalidTagNameLength, "must be at least 2 characters")
	}

	if len(tn) > 8 {
		return domain.NewValidationError(ErrCodeInvalidTagNameLength, "the maximum value is 8 characters")
	}

	return nil
}
