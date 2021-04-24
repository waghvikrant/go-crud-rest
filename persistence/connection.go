package persistence

import (
	logger "github.com/sirupsen/logrus"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func connect() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("data.db"), &gorm.Config{})

	if err != nil {
		logger.Fatal(err)
	}

	logger.Info("Database connected")
	return db
}
