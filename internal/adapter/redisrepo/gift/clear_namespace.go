package cachedGift

import (
	"context"
	"time"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/config"
)

func (r *CacheRepository) ClearNamespace() error {
	c := config.Rds()

	ctx, cancel := context.WithTimeout(context.Background(), c.ContextCancelDuration*time.Second)
	defer cancel()

	cmd := r.db.FlushDB(ctx)
	if cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}
