package repo

import (
	"context"

	be "github.com/syafdia/xaam/internal/domain/entity/business"
)

type BusinessRepo interface {
	FindOneByBusinessID(ctx context.Context, businessID string) (be.GetOneResponse, error)
}
