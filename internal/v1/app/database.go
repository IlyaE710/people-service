package app

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"people/internal/v1/repository/people/entity"
)

func SetupDatabase() (*gorm.DB, error) {
	dsn := "host=localhost user=user password=pass dbname=db port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&entity.People{}, &entity.Country{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
