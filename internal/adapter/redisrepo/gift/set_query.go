package cachedGift

import (
	"context"
	"encoding/json"
	"time"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/config"
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/entity"
)

func (r *GiftRepository) SetQuery(q string, res *entity.PaginatedGiftResponse) error {
	c := config.Rds()

	q = giftQueryNamespace + q

	ctx, cancel := context.WithTimeout(context.Background(), c.ContextCancelDuration*time.Second)
	defer cancel()

	bytes, err := json.Marshal(res)
	if err != nil {
		return err
	}

	cmd := r.db.Set(ctx, q, bytes, c.KeyExpirationDuration*time.Second)
	if cmd.Err() != nil {
		return cmd.Err()
	}

	return nil
}
