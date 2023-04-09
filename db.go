package main

import (
	"fmt"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	dbUser         = os.Getenv("DB_USER")
	dbPwd          = os.Getenv("DB_PASSWORD")
	dbName         = os.Getenv("DB_NAME")
	unixSocketPath = "/cloudsql/thefatladysang:us-central1:test-server"
)

func dbConnection() (*gorm.DB, error) {
	dbURI := fmt.Sprintf("%s:%s@unix(%s)/%s?parseTime=true",
		dbUser, dbPwd, unixSocketPath, dbName)
	db, err := gorm.Open(postgres.Open(dbURI), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&User{}, &Interest{})
	return db, nil
}
