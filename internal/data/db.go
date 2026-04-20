package data

import (
	"github.com/lourencogabe/buzao-bot/internal/config"
	"github.com/lourencogabe/buzao-bot/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var logger *config.Logger

func Connect() {
	logger := config.GetLogger("SQlite")
	db, err := gorm.Open(sqlite.Open("bus.db"), &gorm.Config{})
	if err != nil {
		logger.ErrorF("failed to connect to database: %v", err)
	}

	db.AutoMigrate(&models.BusLine{}, &models.BusTime{})

	DB = db
}
