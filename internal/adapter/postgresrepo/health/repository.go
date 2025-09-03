package health

import "gorm.io/gorm"

type HealthRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *HealthRepository {
	return &HealthRepository{
		db: db,
	}
}
