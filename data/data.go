package data

import (
	"log"
	"os"
	"path/filepath"
)

func GetDataFolder() string {
	roaming, err := os.UserConfigDir()
	if err != nil {
		log.Fatalf("Error getting roaming folder: %v", err)
	}

	dir := filepath.Join(roaming, "Moons", "headpat-osc")

	err = os.MkdirAll(dir, os.ModePerm)
	if err != nil {
		log.Fatalf("Error creating app data folder: %v", err)
	}

	return dir
}
