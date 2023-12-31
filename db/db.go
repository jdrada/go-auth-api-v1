package db

import (
	"github.com/jdrada/go-auth-v1/api/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {
	dsn := "host=**** user=postgres password=***** dbname=**** port=**** sslmode=disable"
	connection, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("could not connect to the database")
	}

	DB = connection

	//AutoMigrate will ONLY create tables, missing columns and missing indexes, and WON’T change existing column’s types or delete unused columns
	connection.AutoMigrate(&model.User{}) 
}
