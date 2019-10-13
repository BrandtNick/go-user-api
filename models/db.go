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

	username := os.Getenv("db_user")
	password := os.Getenv("db_pass")
	dbName := os.Getenv("db_name")
	dbHost := os.Getenv("db_host")

	URI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s", dbHost, username, dbName, password)
	fmt.Println(URI)

	DB, err = gorm.Open("postgres", URI)
	if err != nil {
		fmt.Print(err)
	}
}
