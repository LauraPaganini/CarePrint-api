package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

const endpoint = "careprint-db.cyrujxd0jetc.us-east-2.rds.amazonaws.com"
const user = "root"
const pass = "CarePrint"
const dbName = "careprint-db"

var db *sql.DB

func openDatabaseConnection() {
	var err error
	db, err = sql.Open("mssql",
		"server="+endpoint+";user id="+user+";password="+pass+";")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
}

func closeDatabaseConnection() {
	if db != nil {
		err := db.Close()
		if err != nil {
			log.Fatal(err)
		}
	}
}

func modifyData(statement string) {
	stmt, err := db.Prepare(statement)
	_, err = stmt.Exec()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("done: " + statement)
	}
}
