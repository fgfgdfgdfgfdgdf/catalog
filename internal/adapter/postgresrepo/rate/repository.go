package rate

import (
	"gorm.io/gorm"
)

type RateRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *RateRepository {
	return &RateRepository{
		db: db,
	}
}
