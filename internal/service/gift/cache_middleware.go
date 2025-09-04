package gift

import (
	"net/http"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/entity"
	"github.com/gin-gonic/gin"
)

func (s *Service) CacheMiddleware(c *gin.Context) {
	rawQuery, ok := c.Get("queryParams")
	if !ok {
		c.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	q := rawQuery.(*entity.GiftQuery)

	items, err := s.cacheRepo.GetByQuery(q.RawQuery)
	if err == nil {
		c.AbortWithStatusJSON(http.StatusOK, items)
		return
	}

	c.Next()
}
