package app

import (
	"github.com/raihanmd/fp-superbootcamp-go/helper"
	"github.com/raihanmd/fp-superbootcamp-go/model/entity"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewConnection() *gorm.DB {
	db, err := gorm.Open(postgres.Open(helper.MustGetEnv("DB_DSN")), &gorm.Config{
		QueryFields: true,
		Logger:      logger.Default.LogMode(logger.Info),
	})
	helper.PanicIfError(err)

	err = db.AutoMigrate(&entity.User{}, &entity.Car{}, &entity.CarSpecification{}, &entity.Brand{}, &entity.Review{}, &entity.Comment{}, &entity.Favourite{}, &entity.Profile{})
	helper.PanicIfError(err)

	// create full text index on reviews.title
	db.Exec("CREATE INDEX IF NOT EXISTS idx_title_fulltext ON reviews USING GIN (to_tsvector('english', title))")

	// create index on cars.model
	db.Exec("CREATE EXTENSION IF NOT EXISTS pg_trgm;")
	db.Exec("CREATE INDEX idx_model_gin ON cars USING GIN (model gin_trgm_ops);")

	return db
}
