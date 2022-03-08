package main

import (
	"github.com/gin-gonic/gin"
	"github.com/syafdia/xaam/internal/delivery"
	"github.com/syafdia/xaam/internal/di"
)

func main() {
	appModule := di.GetAppModule()
	repoModule := di.GetRepoModule(appModule)
	useCaseModule := di.GetUseCaseModule(repoModule)

	r := gin.Default()
	r.GET("/ping", delivery.GetPingHandler())
	r.POST("/api/v1/authorisation/check", delivery.PostAuthorisationHandler(useCaseModule))

	r.Run()
}
