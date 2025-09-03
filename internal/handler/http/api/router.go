package api

import (
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/middlewares"
	service_gift "github.com/fgfgdfgdfgfdgdf/catalog/internal/service/gift"
	service_health "github.com/fgfgdfgdfgfdgdf/catalog/internal/service/health"
	service_rate "github.com/fgfgdfgdfgfdgdf/catalog/internal/service/rate"
	"github.com/gin-gonic/gin"
)

func Init(g *gin.Engine, gs *service_gift.Service, rs *service_rate.Service, hs *service_health.Service) {

	g.GET("/gifts", gs.GetGifts, middlewares.RateLimiter)

	g.GET("/healthz", hs.DBHealth)

	admin := g.Group("/admin", middlewares.AdminAccess)
	{
		admin.PUT("/rates", rs.UpdateRates)

		admin.POST("/prices/sync", gs.SyncGiftsPrices)
	}
}
