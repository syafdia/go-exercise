package application

import (
	"context"

	"github.com/syafdia/demo-es/internal/domain"
	"github.com/syafdia/demo-es/internal/domain/report"
	"github.com/syafdia/demo-es/internal/domain/task"
)

type ProjectionHandler struct {
	reportRepository report.ReportRepository
}

func (c *ProjectionHandler) ReceiveEvent(ctx context.Context, event domain.Event) error {
	switch e := event.(type) {
	case task.TaskCreated:
		tr := report.NewTaskReport()

		tr.TaskID = e.ID
		tr.TaskTitle = e.Title
		tr.StartedAt = e.OccuredAt()

		err := c.reportRepository.Store(ctx, tr)
		if err != nil {
			return err
		}

	case task.TaskStatusChanged:
		tr, err := c.reportRepository.FindByTaskID(ctx, e.ID)
		if err != nil {
			return err
		}

		if e.Status != task.TaskStatusCompleted {
			break
		}

		tr.FinishedAt = e.OccuredAt()
		tr.Duration = tr.FinishedAt.Sub(tr.StartedAt)

		err = c.reportRepository.Update(ctx, tr)
		if err != nil {
			return err
		}
	}

	return nil
}
