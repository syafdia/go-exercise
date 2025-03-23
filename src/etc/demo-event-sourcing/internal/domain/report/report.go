package report

import (
	"time"

	"github.com/syafdia/demo-es/internal/domain"
	"github.com/syafdia/demo-es/internal/domain/task"
)

type TaskReport struct {
	ID         domain.ID
	TaskID     domain.ID
	TaskTitle  string
	StartedAt  time.Time
	FinishedAt time.Time
	Duration   time.Duration
}

func NewTaskReport() TaskReport {
	return TaskReport{
		ID: domain.NewID(),
	}
}

func (t *TaskReport) ApplyEvent(event domain.Event) {
	switch e := event.(type) {
	case task.TaskCreated:
		t.TaskID = e.ID
		t.TaskTitle = e.Title
		t.StartedAt = e.OccuredAt()

	case task.TaskStatusChanged:
		if e.Status == task.TaskStatusCompleted {
			t.FinishedAt = e.OccuredAt()
			t.Duration = t.FinishedAt.Sub(t.StartedAt)
		}
	}
}
