package cachedGift

import (
	"github.com/redis/go-redis/v9"
)

type CacheRepository struct {
	db *redis.Client
}

func NewRepository(db *redis.Client) *CacheRepository {
	return &CacheRepository{
		db: db,
	}
}
