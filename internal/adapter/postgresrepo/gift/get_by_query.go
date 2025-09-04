package gift

import (
	"context"
	"math"
	"time"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/config"
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/entity"
)

func (r *GiftRepository) GetByQuery(q *entity.GiftQuery) (*entity.PaginatedGiftResponse, error) {
	pgConf := config.Pg()
	appConf := config.App()

	query := r.db

	giftFilter := &entity.Gift{}

	if q.Type != "" {
		giftFilter.Type = q.Type
	}

	if len(q.Rarity) > 0 {
		query = query.Where("rarity IN ?", q.Rarity)
	}

	if q.MinPriceUsd > 0 {
		query = query.Where("price_usd >= ?", q.MinPriceUsd)
	}
	if q.MaxPriceUsd > 0 {
		query = query.Where("price_usd <= ?", q.MaxPriceUsd)
	}

	if q.Search != "" {
		query = query.Where("name ILIKE ?", q.Search)
	}

	if len(q.Sort.Items) > 0 {
		query = query.Order(q.Sort.Items)
	}

	query = query.Order("id")

	page := max(q.Page, appConf.DefaultPage)

	perPage := q.PerPage
	if perPage <= 0 || perPage > appConf.MaxPerPage {
		perPage = appConf.DefaultPerPage
	}

	offset := (page - 1) * perPage
	query = query.Offset(int(offset)).Limit(int(perPage))

	ctx, cancel := context.WithTimeout(context.Background(), pgConf.ContextCancelDuration*time.Second)
	defer cancel()

	var (
		gifts []*entity.Gift
		total int64
	)

	result := query.WithContext(ctx).
		Model(giftFilter).
		Count(&total).
		Find(&gifts)

	if result.Error != nil {
		return nil, result.Error
	}

	res := &entity.PaginatedGiftResponse{
		Items:     gifts,
		Page:      page,
		Total:     total,
		PerPage:   perPage,
		PageCount: int(math.Round(float64(total) / float64(perPage))),
	}

	return res, nil
}
