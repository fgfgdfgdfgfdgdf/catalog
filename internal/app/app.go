package app

import (
	"fmt"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/adapter/postgresrepo"
	repository_gift "github.com/fgfgdfgdfgfdgdf/catalog/internal/adapter/postgresrepo/gift"
	repository_health "github.com/fgfgdfgdfgfdgdf/catalog/internal/adapter/postgresrepo/health"
	repository_rate "github.com/fgfgdfgdfgfdgdf/catalog/internal/adapter/postgresrepo/rate"
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/adapter/redisrepo"
	cachedGift "github.com/fgfgdfgdfgfdgdf/catalog/internal/adapter/redisrepo/gift"
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/adapter/redisrepo/health"
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/config"
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/handler/http/api"
	service_gift "github.com/fgfgdfgdfgfdgdf/catalog/internal/service/gift"
	service_health "github.com/fgfgdfgdfgfdgdf/catalog/internal/service/health"
	service_rate "github.com/fgfgdfgdfgfdgdf/catalog/internal/service/rate"
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/usecase"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
)

func Run() {
	c := config.Init()

	g := gin.New()

	dbPg := postgresrepo.Init()
	dbRds := redisrepo.Init()

	useCase := InitServices(dbPg, dbRds)
	api.Init(g, useCase)

	g.Run(fmt.Sprintf("%s:%s", c.HOST, c.PORT))
}

func InitServices(dbPg *gorm.DB, dbRds *redis.Client) *usecase.UseCase {
	rateRepo := repository_rate.NewRepository(dbPg)
	giftRepo := repository_gift.NewRepository(dbPg)
	pgHealthRepo := repository_health.NewRepository(dbPg)

	cacheRepo := cachedGift.NewRepository(dbRds)
	rdsHealthRepo := health.NewRepository(dbRds)

	giftService := service_gift.NewService(giftRepo, rateRepo, cacheRepo)
	rateService := service_rate.NewService(rateRepo)
	healthService := service_health.NewService(pgHealthRepo, rdsHealthRepo)

	return &usecase.UseCase{
		RateSvc:   rateService,
		Giftsvc:   giftService,
		HealthSvc: healthService,
	}
}
