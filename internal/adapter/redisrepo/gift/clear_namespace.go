package cachedGift

import (
	"context"
	"time"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/config"
)

func (r *GiftRepository) ClearNamespace() error {
	c := config.Rds()

	ctx, cancel := context.WithTimeout(context.Background(), c.ContextCancelDuration*time.Second)
	defer cancel()

	cmd := r.db.Scan(ctx, 0, giftQueryNamespace, 0)
	if cmd.Err() != nil {
		return cmd.Err()
	}

	iter := cmd.Iterator()

	var keys []string

	for iter.Next(ctx) {
		keys = append(keys, iter.Val())
	}

	intCmd := r.db.Unlink(ctx, keys...)

	if intCmd.Err() != nil {
		return intCmd.Err()
	}

	return nil
}
