package main

import "gorm.io/gorm"

// Database objects model.

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string `gorm:"unique"`
	Movies   []Movie
}

type Movie struct {
	gorm.Model
	Title  string
	Year   int
	Genre  string
	Rating string
	UserID uint
}
