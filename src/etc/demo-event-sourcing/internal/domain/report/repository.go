package report

import (
	"context"

	"github.com/syafdia/demo-es/internal/domain"
)

type ReportRepository interface {
	Store(ctx context.Context, taskReport TaskReport) error
	Update(ctx context.Context, taskReport TaskReport) error
	FindByTaskID(ctx context.Context, taskID domain.ID) (TaskReport, error)
}
