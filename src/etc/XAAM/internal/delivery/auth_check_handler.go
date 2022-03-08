package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/syafdia/xaam/internal/di"
	"github.com/syafdia/xaam/internal/domain/entity"
	ae "github.com/syafdia/xaam/internal/domain/entity/auth"
)

func PostAuthorisationHandler(useCaseModule di.UseCaseModule) gin.HandlerFunc {
	return func(c *gin.Context) {
		request := ae.CheckRequest{}

		err := c.BindJSON(&request)
		if err != nil {
			handleError(c, entity.WrapErr(http.StatusBadRequest, err))
			return
		}

		business, err := useCaseModule.FindOneByBusinessIDUseCase.Execute(c, request.Principal.ID)
		if err != nil {
			handleError(c, err)
			return
		}

		resourceWithActions, err := useCaseModule.FindResourcesByComplianceUseCase.Execute(c, ae.FindResourcesByComplianceRequest{
			IndustryID:    business.IndustryID,
			LegalEntityID: business.LegalEntityID,
			Resources:     []string{request.Resource.Kind},
		})
		if err != nil {
			handleError(c, err)
			return
		}

		response := ae.CheckResponse{}

		for resource, resourceActions := range resourceWithActions {
			actions := map[string]string{}
			for _, ra := range resourceActions {
				actions[ra.Name] = "EFFECT_ALLOWED"
			}

			response.Results = append(response.Results, ae.CheckResultResponse{
				Kind:    resource.Name,
				Actions: actions,
			})
		}

		c.JSON(http.StatusOK, response)
	}
}
