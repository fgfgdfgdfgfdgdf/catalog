package rate

import (
	"net/http"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/entity"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5/pgtype"
)

func (s *Service) UpdateRates(c *gin.Context) {
	var body entity.RateForm

	if err := c.ShouldBind(&body); err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	var UsdPerTon, UsdPerStar pgtype.Numeric

	err := UsdPerTon.Scan(body.UsdPerTon)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	err = UsdPerStar.Scan(body.UsdPerStar)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	newRate := &entity.Rate{
		UsdPerTon:  UsdPerTon,
		UsdPerStar: UsdPerStar,
	}

	err = s.rateRepo.UpdateRates(newRate)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": newRate.ID})
}
