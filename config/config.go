package config

import (
	"errors"

	"gorm.io/gorm"
)

var (
	db *gorm.DB
	logger *Logger
)

// Init initializes the database connection
// Can return an error if any occurs during the initialization
func Init() error {
	return errors.New("Fake error")
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
