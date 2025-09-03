package rate

import (
	"context"
	"time"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/config"
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/entity"
)

func (r *RateRepository) UpdateRates(newRate *entity.Rate) error {
	c := config.Pg()

	ctx, cancel := context.WithTimeout(context.Background(), c.CONTEXT_CANCEL_DURATION*time.Second)
	defer cancel()

	tx := r.db.WithContext(ctx).Begin()

	result := tx.Model(&entity.Rate{}).
		Where(&entity.Rate{
			IsActive: true,
		}).
		Updates(&entity.Rate{
			IsActive: false,
		})

	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	result = tx.Create(newRate)
	if result.Error != nil {
		tx.Rollback()
		return result.Error
	}

	tx.Commit()

	return nil
}
