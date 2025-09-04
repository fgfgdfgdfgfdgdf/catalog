package health

import (
	"context"
	"time"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/config"
)

func (r *HealthRepository) CheckStatus() bool {
	c := config.Rds()

	ctx, cancel := context.WithTimeout(context.Background(), c.ContextCancelDuration*time.Second)
	defer cancel()

	cmd := r.db.Ping(ctx)

	return cmd.Err() == nil
}
