package database

import (
	"log"
	"os"
	"path/filepath"

	"github.com/OPTIC7409/CTMP-API/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() error {
	dataDir := "./data"
	if err := os.MkdirAll(dataDir, os.ModePerm); err != nil {
		log.Fatal("Failed to create data directory:", err)
	}

	dbPath := filepath.Join(dataDir, "taskmanagement.db")

	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		return err
	}

	DB = db

	err = DB.AutoMigrate(&models.User{}, &models.Project{}, &models.Task{})
	if err != nil {
		log.Fatal("Failed to auto migrate:", err)
	}

	log.Println("Connected to the local SQLite database.")
	return nil
}
