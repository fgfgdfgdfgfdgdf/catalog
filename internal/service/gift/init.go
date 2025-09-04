package gift

import "github.com/fgfgdfgdfgfdgdf/catalog/internal/entity"

type giftRepository interface {
	GetByQuery(*entity.GiftQuery) (*entity.PaginatedGiftResponse, error)
	UpdatePricesByRate(*entity.Rate) (int64, error)
}

type rateRepository interface {
	GetRatesInfo() (*entity.Rate, error)
}

type cacheRepository interface {
	GetByQuery(string) (*entity.PaginatedGiftResponse, error)
	SetQuery(string, *entity.PaginatedGiftResponse) error
	ClearNamespace() error
}

type Service struct {
	giftRepo  giftRepository
	rateRepo  rateRepository
	cacheRepo cacheRepository
}

func NewService(gr giftRepository, rr rateRepository, cr cacheRepository) *Service {
	return &Service{
		giftRepo:  gr,
		rateRepo:  rr,
		cacheRepo: cr,
	}
}
