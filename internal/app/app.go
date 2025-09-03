package app

import (
	"fmt"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/adapter/postgresrepo"
	repository_gift "github.com/fgfgdfgdfgfdgdf/catalog/internal/adapter/postgresrepo/gift"
	repository_health "github.com/fgfgdfgdfgfdgdf/catalog/internal/adapter/postgresrepo/health"
	repository_rate "github.com/fgfgdfgdfgfdgdf/catalog/internal/adapter/postgresrepo/rate"
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/config"
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/handler/http/api"
	service_gift "github.com/fgfgdfgdfgfdgdf/catalog/internal/service/gift"
	service_health "github.com/fgfgdfgdfgfdgdf/catalog/internal/service/health"
	service_rate "github.com/fgfgdfgdfgfdgdf/catalog/internal/service/rate"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Run() {
	c := config.Init()

	g := gin.New()

	db := postgresrepo.Init()
	gs, rs, hs := InitServices(db)
	api.Init(g, gs, rs, hs)

	g.Run(fmt.Sprintf("%s:%s", c.HOST, c.PORT))
}

func InitServices(db *gorm.DB) (*service_gift.Service, *service_rate.Service, *service_health.Service) {
	rateRepo := repository_rate.NewRepository(db)
	giftRepo := repository_gift.NewRepository(db)
	healthRepo := repository_health.NewRepository(db)

	giftService := service_gift.NewService(giftRepo, rateRepo)
	rateService := service_rate.NewService(rateRepo)
	healthService := service_health.NewService(healthRepo)

	return giftService, rateService, healthService
}
