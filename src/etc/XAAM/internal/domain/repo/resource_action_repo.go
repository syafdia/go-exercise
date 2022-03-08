package repo

import (
	"context"

	e "github.com/syafdia/xaam/internal/domain/entity"
)

type PolicyRepo interface {
	FindMultipleByIDandTargetType(
		ctx context.Context,
		targetID int64,
		targetType e.TargetType,
	) ([]e.Action, error)
}
