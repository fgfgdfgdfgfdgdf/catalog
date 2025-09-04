package cachedGift

import (
	"github.com/redis/go-redis/v9"
)

const (
	giftQueryNamespace = "giftQuery:"
)

type GiftRepository struct {
	db *redis.Client
}

func NewRepository(db *redis.Client) *GiftRepository {
	return &GiftRepository{
		db: db,
	}
}
