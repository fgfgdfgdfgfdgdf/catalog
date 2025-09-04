package gift

import (
	"errors"
	"net/http"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/entity"
	"github.com/gin-gonic/gin"
)

func (s *Service) QueryValidate(c *gin.Context) {
	var q entity.GiftQuery

	err := c.ShouldBindQuery(&q)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	q.RawQuery = c.Request.URL.Query().Encode()

	err = validateQuery(q)
	if err != nil {
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Set("queryParams", &q)

	c.Next()
}

func validateQuery(q entity.GiftQuery) error {
	for _, item := range q.Sort.Items {
		if _, ok := entity.GiftColumns[item.Column.Name]; !ok {
			return errors.New("invalid name: " + item.Column.Name)
		}
	}

	return nil
}
