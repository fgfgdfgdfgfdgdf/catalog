package postgresrepo

import (
	"fmt"

	"github.com/fgfgdfgdfgfdgdf/catalog/internal/config"
	"github.com/fgfgdfgdfgfdgdf/catalog/internal/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init() *gorm.DB {
	c := config.Pg()

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s",
		c.HOST,
		c.USER,
		c.PASSWORD,
		c.NAME,
		c.PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn))
	if err != nil {
		panic(err)
	}

	err = db.AutoMigrate(&entity.Gift{}, &entity.Rate{})
	if err != nil {
		panic(err)
	}

	return db
}
