package database

import (
	"log"

	"github.com/m-oons/headpat-osc/models"
)

func GetHeadpatCount() int64 {
	var count int64
	tx := Connection.Model(&models.Headpat{}).Count(&count)
	if tx.Error != nil {
		log.Fatalf("Error getting headpats count: %v", tx.Error)
	}
	return count
}
