package service_health

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Service) DBHealth(c *gin.Context) {
	status := s.healthRepo.CheckStatus()

	c.JSON(http.StatusOK, status)
}
