package models

import (
	"fmt"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres" // copied from docs
	"github.com/joho/godotenv"
)

// DB - Database object
var DB *gorm.DB
var err error

func init() {
	err = godotenv.Load()
	if err != nil {
		fmt.Print(err)
	}

	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	username := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASS")

	URI := fmt.Sprintf("dbname=%s host=%s user=%s password=%s sslmode=disable", dbName, dbHost, username, password)
	DB, err = gorm.Open("postgres", URI)
	if err != nil {
		fmt.Print(err)
	}
}
