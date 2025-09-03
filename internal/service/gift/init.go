package gift

import "github.com/fgfgdfgdfgfdgdf/catalog/internal/entity"

type giftRepository interface {
	GetByQuery(entity.GiftQuery) (*entity.PaginatedGiftResponse, error)
	UpdatePricesByRate(*entity.Rate) (int64, error)
}

type rateRepository interface {
	GetRatesInfo() (*entity.Rate, error)
}

type Service struct {
	giftRepo giftRepository
	rateRepo rateRepository
}

func NewService(gr giftRepository, rr rateRepository) *Service {
	return &Service{
		giftRepo: gr,
		rateRepo: rr,
	}
}
