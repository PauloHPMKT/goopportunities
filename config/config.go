package config

import (
	"fmt"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

// Init initializes the database connection
// Can return an error if any occurs during the initialization
func Init() error {
	var err error

	db, err = InitializeSqLite()
	if err != nil {
		return fmt.Errorf("error initializing sqlite: %v", err)
	}
	return nil
}

// return the database connection locally
func GetSqLite() *gorm.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
