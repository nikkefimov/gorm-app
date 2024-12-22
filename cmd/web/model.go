package main

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Username string `gorm:"unique"`
	Password string
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
