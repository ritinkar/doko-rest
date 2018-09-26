package main

import (
	// models
	"doko-rest/models"
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB //database

func init() {
	// handle db
	dbString := fmt.Sprintf(
		"host=%v port=%v user=%v dbname=%v password=%v sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_DBNAME"),
		os.Getenv("DB_PWD"))

	fmt.Println(dbString)

	conn, err := gorm.Open("postgres", dbString)

	if err != nil {
		log.Println(err)
		panic("failed to connect to database")
	} else {
		log.Println("Connection to db established ...")
		db = conn
	}

	db.AutoMigrate(&models.Question{}, &models.Answer{})
}

// GetDB : returns a handle to the DB object
func GetDB() *gorm.DB {
	return db
}
