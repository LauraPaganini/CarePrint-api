package main

import (
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
)

const endpoint = "careprint-db.cyrujxd0jetc.us-east-2.rds.amazonaws.com"
const port = "1433"
const user = "root"
const pass = "CarePrint"
const dbName = "careprint-db"

var db *sql.DB

func openDatabaseConnection() {
	var err error
	db, err = sql.Open("mssql",
		//		"server="+endpoint+";user id="+user+";password="+pass+";database="+dbName+";")
		"sqlserver://"+user+":"+pass+"@"+endpoint+":"+port+"?database="+dbName)
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

func modifyData(statement string) error {
	_, err := db.Exec(statement)
	return err
}

func retrieveData(query string) sql.Rows {
	rows, err := db.Query(query)
	if err != nil {
		log.Fatal(err)
	}
	return *rows
}
