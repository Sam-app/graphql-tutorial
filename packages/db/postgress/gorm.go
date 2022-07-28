package database

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var Db *gorm.DB

func InitDB() error {
	// Connect to database
	dsn := "host=localhost user=samuel dbname=atos port=5432 sslmode=disable TimeZone=Europe/London"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("err: %s", err.Error())
	}
	fmt.Println("Succefully connected to db", db)
	Db = db
	//Migrate()
	return nil
}

// func CloseDB() {
// 	return Db.Close()
// }

func Migrate(models ...interface{}) error {
	err := Db.AutoMigrate(models...)
	if err != nil {
		return fmt.Errorf("err: %s", err.Error())
	}
	return nil
}
