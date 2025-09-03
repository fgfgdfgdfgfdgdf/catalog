package rate

import "github.com/fgfgdfgdfgfdgdf/catalog/internal/entity"

type rateRepository interface {
	UpdateRates(*entity.Rate) error
}

type Service struct {
	rateRepo rateRepository
}

func NewService(rr rateRepository) *Service {
	return &Service{
		rateRepo: rr,
	}
}
