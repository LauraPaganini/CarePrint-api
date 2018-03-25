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

func open() {
	var err error
	db, err = sql.Open("mssql",
		"server="+endpoint+";user id="+user+";password="+pass+";")
	if err != nil {
		log.Fatal(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
	}
}

func close() {
	if db != nil {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}
}
