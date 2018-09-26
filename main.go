package main

import (
	"log"
	"net/http"
	"os"

	// models

	// route handlers
	"doko-rest/controllers"

	// middleware for routers
	"github.com/gorilla/handlers"
	// router
	"github.com/gorilla/mux"
	// orm for go

	// dialect for postgres
	_ "github.com/jinzhu/gorm/dialects/postgres"
	// handle dotenv
	"github.com/joho/godotenv"
)

func main() {

	// handle dotenv
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// close the db connection when main ends
	db := GetDB()
	defer db.Close()

	// mux
	r := mux.NewRouter()

	r.HandleFunc("/heart", controllers.HeartBeat)
	r.HandleFunc("/question", controllers.CreateQuestion)

	// Logged router
	loggedRouter := handlers.LoggingHandler(os.Stdout, r)
	log.Fatal(http.ListenAndServe(":8765", loggedRouter))
}
