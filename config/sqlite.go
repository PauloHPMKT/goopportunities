package config

import (
	"os"

	"github.com/PauloHPMKT/goopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSqLite() (*gorm.DB, error) {
	dbPath := "./db/main.db"
	logger := GetLogger("sqlite")
	// check if the database file exists
	createDirectory("./db", dbPath)

	// Create a new connection to the database
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if err != nil {
		logger.Errorf("SqLite opening error: %v", err)
		return nil, err
	}
	// Migrate the schema
	err = db.AutoMigrate(&schemas.Opening{})
	if err != nil {
		logger.Errorf("SqLite migration error: %v", err)
		return nil, err
	}

	return db, nil
}

func createDirectory(db string, path string) (string, error) {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		logger.Info("Database file does not exist, creating a new one")
		err = os.MkdirAll(db, os.ModePerm) // permanente mode

		if err != nil {
			logger.Errorf("Error creating db directory: %v", err)
			return "", err
		}
		file, err := os.Create(path)
		if err != nil {
			logger.Errorf("Error creating db file: %v", err)
			return "", err
		}
		file.Close()
	}
	return path, nil
}
