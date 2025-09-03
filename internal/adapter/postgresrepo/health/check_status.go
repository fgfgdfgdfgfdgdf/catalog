package health

import (
	"context"
	"time"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/config"
)

func (r *HealthRepository) CheckStatus() bool {
	sqlDB, err := r.db.DB()
	if err != nil {
		return false
	}

	c := config.Pg()

	ctx, cancel := context.WithTimeout(context.Background(), c.CONTEXT_CANCEL_DURATION*time.Second)
	defer cancel()

	err = sqlDB.PingContext(ctx)

	return err == nil
}
