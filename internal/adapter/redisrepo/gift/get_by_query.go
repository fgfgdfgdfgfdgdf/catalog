package cachedGift

import (
	"context"
	"encoding/json"
	"time"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/config"
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/entity"
)

func (r *GiftRepository) GetByQuery(q string) (*entity.PaginatedGiftResponse, error) {
	c := config.Rds()

	ctx, cancel := context.WithTimeout(context.Background(), c.ContextCancelDuration*time.Second)
	defer cancel()

	var response entity.PaginatedGiftResponse

	q = giftQueryNamespace + q

	bytes, err := r.db.Get(ctx, q).Bytes()
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return nil, err
	}

	return &response, nil
}
