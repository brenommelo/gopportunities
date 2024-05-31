package config

import (
	"os"

	"github.com/brenommelo/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("SQLite")

	dbPath := "./db/main.db"

	//check if the database file exists
	_, err := os.Stat(dbPath)
	if os.IsNotExist(err) {
		logger.Info("Data base file not fount")

		//create the database file and directory
		err = os.MkdirAll("./db", os.ModePerm)
		if err != nil {
			return nil, err
		}

		file, err := os.Create(dbPath)
		if err != nil {
			return nil, err
		}
		file.Close()
	}

	//create database and connect
	db, err := gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("Sqlite openming error: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Oppening{})

	if err != nil {
		logger.Errorf("Sqlite automigration error: %v", err)
		return nil, err
	}

	return db, nil
}
