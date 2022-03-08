package source

import (
	"context"

	"github.com/syafdia/xaam/internal/domain/entity"
	be "github.com/syafdia/xaam/internal/domain/entity/business"
	"github.com/syafdia/xaam/internal/domain/repo"
)

type businessRepo struct {
}

var businessIDWithBusiness = map[string]be.GetOneResponse{
	"abcd-0001-56789-asdf": {
		BusinessID:    "abcd-0001-56789-asdf",
		CountryID:     "ID",
		IndustryID:    1,
		LegalEntityID: 1,
	},
	"abcd-0002-56789-asdf": {
		BusinessID:    "abcd-0002-56789-asdf",
		CountryID:     "ID",
		IndustryID:    2,
		LegalEntityID: 2,
	},
}

func NewBusinessRepo() repo.BusinessRepo {
	return &businessRepo{}
}

func (b *businessRepo) FindOneByBusinessID(ctx context.Context, businessID string) (be.GetOneResponse, error) {
	resp, ok := businessIDWithBusiness[businessID]
	if !ok {
		return be.GetOneResponse{}, entity.ErrNotFound
	}

	return resp, nil
}
