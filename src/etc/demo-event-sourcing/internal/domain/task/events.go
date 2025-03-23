package task

import (
	"time"

	"github.com/syafdia/demo-es/internal/domain"
)

type TaskCreated struct {
	ID          domain.ID `json:"-"`
	Timestamp   time.Time `json:"-"`
	Title       string
	Description string
	Status      TaskStatus
}

func (t TaskCreated) AggregateID() domain.ID {
	return t.ID
}

func (t TaskCreated) OccuredAt() time.Time {
	return t.Timestamp
}

type TaskStatusChanged struct {
	ID        domain.ID `json:"-"`
	Timestamp time.Time `json:"-"`
	Status    TaskStatus
}

func (t TaskStatusChanged) AggregateID() domain.ID {
	return t.ID
}

func (t TaskStatusChanged) OccuredAt() time.Time {
	return t.Timestamp
}

type TaskTagAdded struct {
	ID        domain.ID `json:"-"`
	Timestamp time.Time `json:"-"`
	Tag       Tag
}

func (t TaskTagAdded) AggregateID() domain.ID {
	return t.ID
}

func (t TaskTagAdded) OccuredAt() time.Time {
	return t.Timestamp
}

type TaskUpvoted struct {
	ID        domain.ID `json:"-"`
	Timestamp time.Time `json:"-"`
}

func (t TaskUpvoted) AggregateID() domain.ID {
	return t.ID
}

func (t TaskUpvoted) OccuredAt() time.Time {
	return t.Timestamp
}

type TaskDownvoted struct {
	ID        domain.ID `json:"-"`
	Timestamp time.Time `json:"-"`
}

func (t TaskDownvoted) AggregateID() domain.ID {
	return t.ID
}

func (t TaskDownvoted) OccuredAt() time.Time {
	return t.Timestamp
}
