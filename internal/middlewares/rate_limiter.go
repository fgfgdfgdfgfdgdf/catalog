package middlewares

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
)

var instance = cache.New(5*time.Second, 10*time.Minute)

func RateLimiter(c *gin.Context) {
	_, ok := instance.Get(c.ClientIP())
	if ok {
		c.AbortWithStatus(http.StatusTooManyRequests)
		return
	}

	instance.Set(c.ClientIP(), nil, 0)

	c.Next()
}
