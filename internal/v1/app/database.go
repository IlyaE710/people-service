package app

import (
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
	"people/internal/v1/repository/people/entity"
)

func SetupDatabase() (*gorm.DB, error) {
	dbHost := os.Getenv("DB_HOST")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbPort := os.Getenv("DB_PORT")

	dsn := "host=" + dbHost + " user=" + dbUser + " password=" + dbPassword + " dbname=" + dbName + " port=" + dbPort + " sslmode=disable TimeZone=Asia/Shanghai"
	logrus.Infof("Connecting to database: %s", dsn)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		logrus.Errorf("Failed to connect to database: %v", err)
		return nil, err
	}

	logrus.Info("Database connection established")

	err = db.AutoMigrate(&entity.People{}, &entity.Country{})
	if err != nil {
		logrus.Errorf("Failed to perform database auto migration: %v", err)
		return nil, err
	}

	logrus.Info("Database auto migration completed")

	return db, nil
}
