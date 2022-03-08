package delivery

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/syafdia/xaam/internal/domain/entity"
)

func handleError(ctx *gin.Context, err error) {
	if errr, ok := err.(*entity.Err); ok {
		ctx.JSON(errr.Status, gin.H{
			"message": errr.Error(),
		})
		return
	}

	ctx.JSON(http.StatusInternalServerError, gin.H{
		"message": err.Error(),
	})
}
