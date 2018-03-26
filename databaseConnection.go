package main

import (
	"database/sql"
	"fmt"
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

func modifyData(statement string) {
	//stmt, err := db.Prepare(statement)
	//if err != nil {
	//	fmt.Println(err)
	//} else {
	//	fmt.Println("prepped: " + statement)
	//}
	_, err := db.Exec(statement)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("executed: " + statement)
	}

//	defer rows.Close()
//	fmt.Println(rows)
//	for rows.Next() {
//		var (
//			email string
//			pass string
//		)

       // 	if err := rows.Scan(&email, &pass); err != nil {
     //           	log.Fatal(err)
   //     	}
 //       	fmt.Println(email + " " + pass)
//	}
}
