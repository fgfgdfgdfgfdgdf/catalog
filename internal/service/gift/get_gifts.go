package gift

import (
	"fmt"
	"net/http"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/entity"
	"github.com/gin-gonic/gin"
)

func (s *Service) GetGifts(c *gin.Context) {
	rawQuery, ok := c.Get(entity.ContextQueryKey)
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	q := rawQuery.(*entity.GiftQuery)

	response, err := s.giftRepo.GetByQuery(q)
	if err != nil {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	err = s.cacheRepo.SetQuery(q.RawQuery, response)
	if err != nil {
		fmt.Println(err.Error())
	}

	c.JSON(http.StatusOK, response)
}
