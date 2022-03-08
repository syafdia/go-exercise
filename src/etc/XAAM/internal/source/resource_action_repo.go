package source

import (
	"context"

	"github.com/jmoiron/sqlx"
	e "github.com/syafdia/xaam/internal/domain/entity"
	"github.com/syafdia/xaam/internal/domain/repo"
)

const (
	sqlFindMultipleByIDandTargetType = `
		SELECT
			ra.id,
			ra.resource_id,
			ra.name
		FROM target_resource_actions tra 
		JOIN resource_actions ra ON ra.id = tra.resource_action_id 
		WHERE tra.target_id = $1 AND tra.target_type = $2;
	`
)

type PolicyRepo struct {
	db *sqlx.DB
}

func NewPolicyRepo(db *sqlx.DB) repo.PolicyRepo {
	return &PolicyRepo{db: db}
}

func (r *PolicyRepo) FindMultipleByIDandTargetType(
	ctx context.Context,
	targetID int64,
	targetType e.TargetType,
) ([]e.Action, error) {
	results := []e.Action{}
	err := r.db.SelectContext(ctx, &results, sqlFindMultipleByIDandTargetType, targetID, targetType)
	if err != nil {
		return nil, err
	}

	return results, nil
}
