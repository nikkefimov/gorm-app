package main

import (
	"regexp"

	"gorm.io/gorm"
)

// User represents the user account in the database
type User struct {
	gorm.Model
	//UserID   uint   `gorm:"not null"`
	Username string `gorm:"unique;not null;size:15"`
	Password string `gorm:"not null"`
}

// Validate checks if the username meets the requirements
func (u *User) Validate() bool {
	// Username must be 1-15 letters only
	usernameRegex := regexp.MustCompile(`^[A-Za-z]{1,15}$`)
	return usernameRegex.MatchString(u.Username)
}

// Movie represents the movie entry in the database
type Movie struct {
	gorm.Model
	UserID uint   `gorm:"not null"`
	User   User   `gorm:"foreignKey:UserID"`
	Title  string `gorm:"not null;size:25"`
	Year   int    `gorm:"not null"`
	Genre  string `gorm:"not null;size:100"`
	Rating int    `gorm:"not null"`
}

// Validate checks if the movie meets the requirements
func (m *Movie) Validate() bool {
	// Title: 1-25 characters
	// Year: between 1950 and current year
	// Rating: 1-10
	currentYear := 2025 // As specified in requirements
	return len(m.Title) > 0 && len(m.Title) <= 25 &&
		m.Year >= 1950 && m.Year <= currentYear &&
		m.Rating >= 1 && m.Rating <= 10
}
