package usecase

import "github.com/gin-gonic/gin"

type giftService interface {
	GetGifts(c *gin.Context)
	ValidateQueryMiddleware(c *gin.Context)
	SyncGiftsPrices(c *gin.Context)
	CacheMiddleware(c *gin.Context)
}

type rateService interface {
	UpdateRates(c *gin.Context)
}

type healthService interface {
	DBHealth(c *gin.Context)
}

type UseCase struct {
	HealthSvc healthService
	RateSvc   rateService
	Giftsvc   giftService
}
