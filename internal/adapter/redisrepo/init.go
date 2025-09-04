package redisrepo

import (
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/config"
	"github.com/redis/go-redis/v9"
)

func Init() *redis.Client {
	c := config.Rds()

	return redis.NewClient(&redis.Options{
		Addr:     c.HOST,
		Password: c.PASSWORD,
		Username: c.USER,
	})
}
