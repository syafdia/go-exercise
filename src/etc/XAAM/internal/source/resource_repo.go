package source

import (
	"context"

	"github.com/jmoiron/sqlx"
	e "github.com/syafdia/xaam/internal/domain/entity"
	"github.com/syafdia/xaam/internal/domain/repo"
)

const (
	sqlFindMultipleByIDs = `
		SELECT
			r.id,
			r.name
		FROM resources AS r 
		WHERE r.id IN (?);
	`

	sqlFindMultipleByNames = `
		SELECT
			r.id,
			r.name
		FROM resources AS r 
		WHERE r.name IN (?);
	`
)

type resourceRepo struct {
	db *sqlx.DB
}

func NewResourceRepo(db *sqlx.DB) repo.ResourceRepo {
	return &resourceRepo{db: db}
}

func (r *resourceRepo) FindMultipleByIDs(ctx context.Context, ids []int64) (map[int64]e.Resource, error) {
	query, args, err := sqlx.In(sqlFindMultipleByIDs, ids)
	if err != nil {
		return nil, err
	}

	query = r.db.Rebind(query)
	results := []e.Resource{}

	err = r.db.SelectContext(ctx, &results, query, args...)
	if err != nil {
		return nil, err
	}

	idWithResource := map[int64]e.Resource{}
	for _, result := range results {
		idWithResource[result.ID] = result
	}

	return idWithResource, nil
}

func (r *resourceRepo) FindMultipleByNames(ctx context.Context, names []string) (map[string]e.Resource, error) {
	query, args, err := sqlx.In(sqlFindMultipleByNames, names)
	if err != nil {
		return nil, err
	}

	query = r.db.Rebind(query)
	results := []e.Resource{}

	err = r.db.SelectContext(ctx, &results, query, args...)
	if err != nil {
		return nil, err
	}

	nameWithResource := map[string]e.Resource{}
	for _, result := range results {
		nameWithResource[result.Name] = result
	}

	return nameWithResource, nil
}
