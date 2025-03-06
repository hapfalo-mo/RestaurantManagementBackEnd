package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func Connect() {
	db_driver := "mysql"
	db_password := "HJ10xugb123*"
	db_user := "root"
	db_host := "localhost"
	db_name := "restaurantDB"
	// Connecting to the database
	db, err := sql.Open(db_driver, db_user+":"+db_password+"@tcp("+db_host+")/"+db_name)
	if err != nil {
		log.Fatal(err)
	}
	// Use Ping() to check if the databas is connected
	err = db.Ping()
	if err != nil {
		fmt.Println("Something wrong with connection", err)
		log.Fatal(err)
	}
	fmt.Println("Connected to the database")
	DB = db
}
