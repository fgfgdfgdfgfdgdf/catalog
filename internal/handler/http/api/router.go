package api

import (
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/middlewares"
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/usecase"
	"github.com/gin-gonic/gin"
)

func Init(g *gin.Engine, useCase *usecase.UseCase) {

	g.GET("/gifts",
		useCase.Giftsvc.GetGifts,
		useCase.Giftsvc.CacheMiddleware,
		useCase.Giftsvc.ValidateQueryMiddleware,
		middlewares.RateLimiter,
	)

	g.GET("/healthz", useCase.HealthSvc.DBHealth)

	admin := g.Group("/admin")
	{
		admin.PUT("/rates", useCase.RateSvc.UpdateRates)

		admin.POST("/prices/sync", useCase.Giftsvc.SyncGiftsPrices)
	}
}
