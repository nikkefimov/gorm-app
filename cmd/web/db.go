package main

import (
	"log"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connection to MySQL database.

var DB *gorm.DB

func ConnectDB() {
	dsn := os.Getenv("DSN") // Update DNS with your database details for successful connect.
	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error to connect DB: %v", err)
	}
}
