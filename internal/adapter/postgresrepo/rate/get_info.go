package rate

import (
	"context"
	"time"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/config"
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/entity"
)

func (r *RateRepository) GetRatesInfo() (*entity.Rate, error) {
	c := config.Pg()

	ctx, cancel := context.WithTimeout(context.Background(), c.ContextCancelDuration*time.Second)
	defer cancel()

	var rate entity.Rate

	result := r.db.WithContext(ctx).
		Where(&entity.Rate{
			IsActive: true,
		}).
		First(&rate)

	if result.Error != nil {
		return nil, result.Error
	}

	return &rate, nil

}
