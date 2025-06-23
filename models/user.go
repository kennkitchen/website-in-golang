package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"os"
	"strconv"
)

var dbConn DBConn

type DBConn struct {
	DbName string
	DbUser string
	DbPass string
	DbHost string
	DbPort int
}

func init() {
	err := godotenv.Load()
	if err != nil {
		//log.Fatal
		println("Error loading .env file")
	}

	dbConn.DbName = os.Getenv("DBNAME")
	dbConn.DbUser = os.Getenv("DBUSER")
	dbConn.DbPass = os.Getenv("DBPASS")
	dbConn.DbHost = os.Getenv("DBHOST")
	dbConn.DbPort, _ = strconv.Atoi(os.Getenv("DBPORT"))
}

func InsertUserRow(email string, password string) (int, error) {
	conn, err := dbConnect()
	defer conn.Close()

	_, err = conn.Exec("INSERT INTO users (email, password, created) VALUES (?, ?, NOW())", email, password)
	return 0, err
}

func dbConnect() (*sql.DB, error) {
	db, err := sql.Open("mysql", os.Getenv("DBUSER")+":"+os.Getenv("DBPASS")+"@tcp("+os.Getenv("DBHOST")+":"+os.Getenv("DBPORT")+")/"+os.Getenv("DBNAME")+"?tls=skip-verify&autocommit=true")
	//if err != nil {
	//	log.Fatal(err)
	//}

	return db, err
}
