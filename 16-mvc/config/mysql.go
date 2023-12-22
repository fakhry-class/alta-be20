package config

import (
	"fakhry/mvc/models"
	"fmt"
	"os"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// database connection
func InitDB() {
	// declare struct config & variable connectionString
	DBConfig := os.Getenv("CONNECTION_DB")
	connectionString := fmt.Sprintf("%s?charset=utf8&parseTime=True&loc=Local", DBConfig)
	var err error
	DB, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})

	if err != nil {
		panic(err)
	}
}

// db migration
func InitialMigration() {
	DB.AutoMigrate(&models.User{})
	/*
		TODO 2
		migrate struct product
	*/
}
