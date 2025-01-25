package model

import "gorm.io/gorm"

// Database objects model.

type User struct {
	gorm.Model
	ID       uint   `gorm:"primaryKey"`
	Username string `gorm:"unique;not null"`
	Password string `gorm:"not null"`
	Movies   []Movie
}

type Movie struct {
	gorm.Model
	User   User   `form:"not null"`
	ID     uint   `gorm:"primaryKey"`
	Title  string `gorm:"unique;not null"`
	Year   int    `gorm:"not null"`
	Genre  string `gorm:"not null"`
	Rating string `gorm:"not null"`
	UserID uint   `gorm:"not null"`
}
