package repo

import (
	"context"

	e "github.com/syafdia/xaam/internal/domain/entity"
)

// mockgen -source=internal/domain/repo/resource_repo.go -destination=internal/mock/domain/repo/resource_repo.go
type ResourceRepo interface {
	FindMultipleByIDs(ctx context.Context, ids []int64) (map[int64]e.Resource, error)
	FindMultipleByNames(ctx context.Context, names []string) (map[string]e.Resource, error)
}
