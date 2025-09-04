package redisrepo

import (
	"fmt"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/config"
	"github.com/redis/go-redis/v9"
)

func Init() *redis.Client {
	c := config.Rds()

	return redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", c.HOST, c.PORT),
		Password: c.PASSWORD,
		Username: c.USER,
	})
}
