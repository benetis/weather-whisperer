package storage

import (
	"github.com/benetis/weather-whisperer/internal/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
)

func InitDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data/forecast.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database")
	}

	err = db.AutoMigrate(&entities.Place{}, &entities.Forecast{})
	if err != nil {
		log.Fatal("failed to migrate database")
	}

	return db
}
