package main

import (
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// Connection to MySQL database.

var db *gorm.DB

func ConnectDB() {
	dsn := "" // Update DSN with your database details for successful connect.
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error to connect DB: %v", err)
	}

	/*// Migrate DB.
	err = DB.AutoMigrate(&user{}, &movie{})
	if err != nil {
		log.Fatal("Failed to migrate database", err)
	}*/
	log.Println("Database initialized successfully")
}
