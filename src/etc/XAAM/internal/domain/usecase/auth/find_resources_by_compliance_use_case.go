package auth

import (
	"context"
	"net/http"

	e "github.com/syafdia/xaam/internal/domain/entity"
	ae "github.com/syafdia/xaam/internal/domain/entity/auth"
	"github.com/syafdia/xaam/internal/domain/repo"
	"github.com/syafdia/xaam/pkg/slices"
)

type FindResourcesByComplianceUseCase interface {
	Execute(ctx context.Context, request ae.FindResourcesByComplianceRequest) (map[e.Resource][]e.Action, error)
}

type findResourcesByComplianceUC struct {
	PolicyRepo   repo.PolicyRepo
	ResourceRepo repo.ResourceRepo
}

func NewFindResourcesByComplianceUseCase(
	PolicyRepo repo.PolicyRepo,
	resourceRepo repo.ResourceRepo,
) FindResourcesByComplianceUseCase {
	return &findResourcesByComplianceUC{
		PolicyRepo:   PolicyRepo,
		ResourceRepo: resourceRepo,
	}
}

func (f *findResourcesByComplianceUC) Execute(
	ctx context.Context,
	request ae.FindResourcesByComplianceRequest,
) (map[e.Resource][]e.Action, error) {
	actionIndustries, err := f.PolicyRepo.FindMultipleByIDandTargetType(ctx, request.IndustryID, e.TargetTypeIndustry)
	if err != nil {
		return nil, e.WrapErr(http.StatusInternalServerError, err)
	}

	actionLegalEntities, err := f.PolicyRepo.FindMultipleByIDandTargetType(ctx, request.LegalEntityID, e.TargetTypeLegalEntity)
	if err != nil {
		return nil, e.WrapErr(http.StatusInternalServerError, err)
	}

	actionIntersections := f.getIntersections(actionIndustries, actionLegalEntities)
	resourceIDWithActions := f.groupActionsByResourceID(actionIntersections)

	resourceIDs := []int64{}
	for resourceID := range resourceIDWithActions {
		resourceIDs = append(resourceIDs, resourceID)
	}

	idWithResource, err := f.ResourceRepo.FindMultipleByIDs(ctx, resourceIDs)
	if err != nil {
		return nil, e.WrapErr(http.StatusInternalServerError, err)
	}

	resourceWithActions := map[e.Resource][]e.Action{}
	for resourceID, actions := range resourceIDWithActions {
		resource, ok := idWithResource[resourceID]
		if !ok {
			continue
		}

		if !slices.ContainsStr(request.Resources, resource.Name) {
			continue
		}

		resourceWithActions[resource] = actions
	}

	return resourceWithActions, nil
}

func (f *findResourcesByComplianceUC) getIntersections(xs []e.Action, ys []e.Action) []e.Action {
	actionWithTotal := map[e.Action]int{}

	for _, x := range xs {
		actionWithTotal[x] = actionWithTotal[x] + 1
	}

	for _, y := range ys {
		actionWithTotal[y] = actionWithTotal[y] + 1
	}

	results := []e.Action{}
	for action, total := range actionWithTotal {
		if total == 1 {
			continue
		}

		results = append(results, action)
	}

	return results
}

func (f *findResourcesByComplianceUC) groupActionsByResourceID(actions []e.Action) map[int64][]e.Action {
	resourceIDWithActions := map[int64][]e.Action{}

	for _, action := range actions {
		_, ok := resourceIDWithActions[action.ResourceID]
		if !ok {
			resourceIDWithActions[action.ResourceID] = []e.Action{}
		}

		resourceIDWithActions[action.ResourceID] = append(resourceIDWithActions[action.ResourceID], action)
	}

	return resourceIDWithActions
}
