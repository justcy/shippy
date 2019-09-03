package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"os"
)

func CreateConnection() (*gorm.DB, error) {

	// Get database details from environment variables
	host := os.Getenv("POSTGRES_HOST")
	user := os.Getenv("POSTGRES_USER")
	DBName := os.Getenv("POSTGRES_NAME")
	password := os.Getenv("POSTGRES_PASSWORD")

	return gorm.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s user=%s dbname=%s sslmode=disable password=%s",
			host, user, DBName, password,
		),
	)
}
