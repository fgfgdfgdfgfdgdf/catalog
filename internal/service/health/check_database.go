package service_health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Service) DBHealth(c *gin.Context) {
	pgStatus := s.postgresRepo.CheckStatus()
	rdsStatus := s.redisRepo.CheckStatus()

	c.JSON(http.StatusOK, gin.H{
		"postgres": gin.H{
			"ok": pgStatus,
		},
		"redis": gin.H{
			"ok": rdsStatus,
		},
	},
	)
}
