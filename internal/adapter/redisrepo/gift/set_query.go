package cachedGift

import (
	"context"
	"time"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/config"
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/entity"
)

func (r *CacheRepository) SetQuery(q string, res *entity.PaginatedGiftResponse) error {
	c := config.Rds()

	ctx, cancel := context.WithTimeout(context.Background(), c.ContextCancelDuration*time.Second)
	defer cancel()

	cmd := r.db.Set(ctx, q, res, c.KeyExpirationDuration*time.Second)
	if cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}
