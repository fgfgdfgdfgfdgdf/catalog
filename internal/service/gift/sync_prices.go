package gift

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Service) SyncGiftsPrices(c *gin.Context) {
	rate, err := s.rateRepo.GetRatesInfo()
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	updated, err := s.giftRepo.UpdatePricesByRate(rate)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = s.cacheRepo.ClearNamespace()
	if err != nil {

	}

	c.JSON(http.StatusOK, gin.H{"updated": updated})
}
