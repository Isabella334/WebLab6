package db

import (
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	DB, err := gorm.Open(sqlite.Open("series.db"), &gorm.Config{})
	if err != nil {
    log.Fatal("Failed to connect to DB", err)
  }
}
