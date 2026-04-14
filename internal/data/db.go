package data

import (
	"log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"github.com/lourencogabe/buzao-bot/internal/models"
)

var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(sqlite.Open("bus.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Erro ao conectar no banco:", err)
	}

	db.AutoMigrate(&models.BusLine{})

	DB = db
}
