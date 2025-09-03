package middlewares

import (
	"github.com/gin-gonic/gin"
)

func AdminAccess(c *gin.Context) {
	c.Next()
}
