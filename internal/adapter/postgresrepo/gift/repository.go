package gift

import (
	"gorm.io/gorm"
)

type GiftRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *GiftRepository {
	return &GiftRepository{
		db: db,
	}
}
