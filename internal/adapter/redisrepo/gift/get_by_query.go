package cachedGift

import (
	"context"
	"time"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/config"
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/entity"
)

func (r *GiftRepository) GetByQuery(q string) (*entity.PaginatedGiftResponse, error) {
	c := config.Rds()

	ctx, cancel := context.WithTimeout(context.Background(), c.ContextCancelDuration*time.Second)
	defer cancel()

	var response entity.PaginatedGiftResponse

	cmd := r.db.HGetAll(ctx, q)
	if cmd.Err() != nil {
		return nil, cmd.Err()
	}

	err := cmd.Scan(&response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
