package database

import (
	"log"
	"time"

	"github.com/m-oons/headpat-osc/models"
)

func CreateHeadpat() bool {
	headpat := models.Headpat{
		Time: time.Now(),
	}

	tx := Connection.Create(&headpat)
	if tx.Error != nil {
		log.Printf("Error creating headpat: %v", tx.Error)
		return false
	}

	return true
}

func GetHeadpatCount() int64 {
	var count int64
	tx := Connection.Model(&models.Headpat{}).Count(&count)
	if tx.Error != nil {
		log.Printf("Error getting headpats count: %v", tx.Error)
	}
	return count
}
