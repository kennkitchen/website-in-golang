package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
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
	//db, err := sql.Open("mysql", os.Getenv("DBUSER")+":"+os.Getenv("DBPASS")+"@tcp("+os.Getenv("DBHOST")+":"+os.Getenv("DBPORT")+")/"+os.Getenv("DBNAME")+"?tls=skip-verify&autocommit=true")
	//if err != nil {
	//	log.Fatal(err)
	//}
	//defer db.Close()
	//
	//log.Info("Connected!")

	config.RunServer()
}
