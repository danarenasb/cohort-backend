package main

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
// dbUser                 = os.Getenv("DB_USER")
// dbPwd                  = os.Getenv("DB_PASS")
// instanceConnectionName = os.Getenv("INSTANCE_CONNECTION_NAME")
// dbName                 = os.Getenv("DB_NAME")
// socketDir              = "/cloudsql"

)

func dbConnection() (*gorm.DB, error) {
	// dsn := fmt.Sprintf("user=%s password=%s database=%s host=%s/%s", dbUser, dbPwd, dbName, socketDir, instanceConnectionName)
	dsn := "host=localhost user=postgres password=postgres dbname=postgres port=5432"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(&Associate{}, &TechCheck{}, &Score{})
	return db, nil
}
