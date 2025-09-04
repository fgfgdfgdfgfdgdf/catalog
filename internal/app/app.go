package app

import (
	"fmt"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/adapter/postgresrepo"
	repository_gift "github.com/fgfgdfgdfgfdgdf/catalog/internal/adapter/postgresrepo/gift"
	repository_health "github.com/fgfgdfgdfgfdgdf/catalog/internal/adapter/postgresrepo/health"
	repository_rate "github.com/fgfgdfgdfgfdgdf/catalog/internal/adapter/postgresrepo/rate"
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/adapter/redisrepo"
	cachedGift "github.com/fgfgdfgdfgfdgdf/catalog/internal/adapter/redisrepo/gift"
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/config"
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/handler/http/api"
	service_gift "github.com/fgfgdfgdfgfdgdf/catalog/internal/service/gift"
	service_health "github.com/fgfgdfgdfgfdgdf/catalog/internal/service/health"
	service_rate "github.com/fgfgdfgdfgfdgdf/catalog/internal/service/rate"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Run() {
	c := config.Init()

	g := gin.New()

	dbPg := postgresrepo.Init()
	dbRds := redisrepo.Init()

	gs, rs, hs := InitServices(dbPg, dbRds)
	api.Init(g, gs, rs, hs)

	g.Run(fmt.Sprintf("%s:%s", c.HOST, c.PORT))
}

func InitServices(dbPg *gorm.DB, dbRds *redis.Client) (*service_gift.Service, *service_rate.Service, *service_health.Service) {
	rateRepo := repository_rate.NewRepository(dbPg)
	giftRepo := repository_gift.NewRepository(dbPg)
	healthRepo := repository_health.NewRepository(dbPg)
	cacheRepo := cachedGift.NewRepository(dbRds)

	giftService := service_gift.NewService(giftRepo, rateRepo, cacheRepo)
	rateService := service_rate.NewService(rateRepo)
	healthService := service_health.NewService(healthRepo)

	return giftService, rateService, healthService
}
