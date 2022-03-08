package business

import (
	"context"

	be "github.com/syafdia/xaam/internal/domain/entity/business"
	"github.com/syafdia/xaam/internal/domain/repo"
)

type FindOneByBusinessIDUseCase interface {
	Execute(ctx context.Context, businessID string) (be.GetOneResponse, error)
}

type findOneByBusinessIDUC struct {
	businessRepo repo.BusinessRepo
}

func NewFindOneByBusinessIDUseCase(businessRepo repo.BusinessRepo) FindOneByBusinessIDUseCase {
	return &findOneByBusinessIDUC{businessRepo: businessRepo}
}

func (g *findOneByBusinessIDUC) Execute(ctx context.Context, businessID string) (be.GetOneResponse, error) {
	return g.businessRepo.FindOneByBusinessID(ctx, businessID)
}
