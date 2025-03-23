package task

import (
	"context"

	"github.com/syafdia/demo-es/internal/domain"
)

type TaskRepository interface {
	Store(ctx context.Context, task Task) error
	Find(ctx context.Context, id domain.ID) (Task, error)
}
