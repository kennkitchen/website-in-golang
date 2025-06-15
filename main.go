package main

import (
	"database/sql"
	"fmt"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"kennkitchen/config"
	"os"
)

func main() {
	file, err := os.OpenFile("logs/webserver.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(file)

	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// TODO remove
	// testing DB connection
	psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", "localhost", 5432, "kennusr", "secret", "kenndb")

	// open database
	db, err := sql.Open("postgres", psqlconn)
	if err != nil {
		log.Fatal("Unabled to connect to database")
	}

	// close database
	defer db.Close()

	// check db
	err = db.Ping()
	if err != nil {
		log.Fatal("No response from database")
	}

	log.Info("Connected!")
	// end test DB connection

	config.RunServer()
}
