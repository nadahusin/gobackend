package orm

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	host := os.Getenv("DB_HOST")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s", host, user, password, dbName)

	gormDb, err := gorm.Open(postgres.Open(config), &gorm.Config{})
	if err != nil {
		return nil, errors.New("failed to initiate database")
	}

	db, err := gormDb.DB()
	if err != nil {
		return nil, errors.New("failed to initiate database")
	}

	db.SetConnMaxIdleTime(10)
	db.SetMaxOpenConns(100)

	return gormDb, nil
}
