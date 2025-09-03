package gift

import (
	"context"
	"time"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/config"
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/entity"
	"gorm.io/gorm"
)

func (r *GiftRepository) UpdatePricesByRate(rate *entity.Rate) (int64, error) {

	c := config.Pg()

	ctx, cancel := context.WithTimeout(context.Background(), c.CONTEXT_CANCEL_DURATION*time.Second)
	defer cancel()

	result := r.db.WithContext(ctx).
		Session(&gorm.Session{AllowGlobalUpdate: true}).
		Model(&entity.Gift{}).
		Updates(map[string]any{
			"price_stars": gorm.Expr("ROUND( price_usd / ?::numeric, 0)", rate.UsdPerStar),
			"price_ton":   gorm.Expr("ROUND( price_usd / ?::numeric, 2)", rate.UsdPerTon),
		})

	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
