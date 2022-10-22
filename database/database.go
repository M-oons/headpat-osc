package database

import (
	"fmt"
	"log"

	"github.com/m-oons/headpat-osc/data"
	"github.com/m-oons/headpat-osc/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Connection *gorm.DB

func SetupDatabase() {
	dir := data.GetDataFolder()
	dbPath := fmt.Sprintf("%s/headpats.db", dir)

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	err = db.AutoMigrate(&models.Headpat{})
	if err != nil {
		log.Fatalf("Error migrating Headpat model: %v", err)
	}

	Connection = db
}
