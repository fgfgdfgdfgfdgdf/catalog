package health

import "github.com/redis/go-redis/v9"

type HealthRepository struct {
	db *redis.Client
}

func NewRepository(db *redis.Client) *HealthRepository {
	return &HealthRepository{
		db: db,
	}
}
