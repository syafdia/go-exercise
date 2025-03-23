package task

import (
	"errors"
	"time"

	"github.com/syafdia/demo-es/internal/domain"
)

type TaskStatus string

const (
	TaskStatusNotStarted TaskStatus = "NOT_STARTED"
	TaskStatusInProgress TaskStatus = "IN_PROGRESS"
	TaskStatusInReview   TaskStatus = "IN_REVIEW"
	TaskStatusCompleted  TaskStatus = "COMPLETED"
)

var (
	ErrTaskTitleTooShort              = errors.New("task: title is too shorts")
	ErrTaskStatusTransitionNotAllowed = errors.New("task: status transition is not allowed")
	ErrTaskCanNotContainsMoreTag      = errors.New("task: can't add more tag")
	ErrTaskDiscussionContentTooShort  = errors.New("task: content is too shorts")
	ErrTaskCanNotVoteCompletedTask    = errors.New("task: can not vote completed task")
)

type Tag struct {
	ID   domain.ID
	Name string
}

func NewTag(name string) Tag {
	return Tag{
		ID:   domain.NewID(),
		Name: name,
	}
}

type Task struct {
	ID          domain.ID
	Title       string
	Description string
	Status      TaskStatus
	Tags        []Tag
	TotalVotes  int

	events  []domain.Event
	version int
}

func NewTask(title string, description string) (Task, error) {
	if len(title) < 4 {
		return Task{}, ErrTaskTitleTooShort
	}

	// Implement some business validation here
	// ...
	// ...

	t := Task{}
	t.AddEvent(TaskCreated{
		ID:          domain.NewID(),
		Title:       title,
		Description: description,
		Status:      TaskStatusNotStarted,
		Timestamp:   time.Now(),
	})

	return t, nil
}

func (t *Task) ChangeStatus(newStatus TaskStatus) error {
	if t.Status == TaskStatusNotStarted {
		if newStatus == TaskStatusCompleted {
			return ErrTaskStatusTransitionNotAllowed
		}
	}

	// Implement some business validation here
	// ...
	// ...

	t.AddEvent(TaskStatusChanged{
		ID:        t.ID,
		Status:    newStatus,
		Timestamp: time.Now(),
	})

	return nil
}

func (t *Task) AddTag(tag Tag) error {
	if len(t.Tags) == 4 {
		return ErrTaskCanNotContainsMoreTag
	}

	t.AddEvent(TaskTagAdded{
		ID:        t.ID,
		Tag:       tag,
		Timestamp: time.Now(),
	})

	return nil
}

func (t *Task) Upvote() error {
	if t.Status == TaskStatusCompleted {
		return ErrTaskCanNotVoteCompletedTask
	}

	t.AddEvent(TaskUpvoted{
		ID:        t.ID,
		Timestamp: time.Now(),
	})

	return nil
}

func (t *Task) Downvote() error {
	if t.Status == TaskStatusCompleted {
		return ErrTaskCanNotVoteCompletedTask
	}

	t.AddEvent(TaskDownvoted{
		ID:        t.ID,
		Timestamp: time.Now(),
	})

	return nil
}

func (t *Task) AddEvent(event domain.Event) {
	t.events = append(t.events, event)
	t.ApplyEvent(event)
}

func (t *Task) ApplyEvent(event domain.Event) {
	switch e := event.(type) {
	case TaskCreated:
		t.ID = e.ID
		t.Title = e.Title
		t.Description = e.Description
		t.Status = e.Status

	case TaskStatusChanged:
		t.Status = e.Status

	case TaskTagAdded:
		t.Tags = append(t.Tags, e.Tag)

	case TaskUpvoted:
		t.TotalVotes++

	case TaskDownvoted:
		t.TotalVotes--
	}

	t.version++
}

func (t *Task) GetEvents() []domain.Event {
	return t.events
}
