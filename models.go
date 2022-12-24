package main

import (
	"gorm.io/gorm"
)

type Users struct {
	gorm.Model
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	ZipCode string  `json:"zip_code"`
	Admin   bool    `json:"admin"`
	Prizes  []Prize `gorm:"many2many:user_prizes;"`
}

type Prize struct {
	gorm.Model
	Name     string `json:"name"`
	Quantity int    `json:"quantity"`
}
