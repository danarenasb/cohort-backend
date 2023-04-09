package main

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name           string         `json:"name"`
	Email          string         `json:"email"`
	Password       string         `json:"password"`
	ZipCode        string         `json:"zip_code"`
	Dob            string         `json:"dob"`
	Admin          bool           `json:"admin"`
	Interests      []Interest     `gorm:"many2many:user_interests";json:"interests"`
	LogintAttempts []LoginAttempt `json:"login_attempts"`
}

type LoginAttempt struct {
	gorm.Model
	Timestamp time.Time `json:"timestamp"`
	UserID    uint      `json:"user_id"`
}

type Interest struct {
	gorm.Model
	Name string `json:"name"`
}
