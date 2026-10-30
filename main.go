package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host         = "localhost"
	port         = 5432
	databaseName = "test_db"
	username     = "root"
	password     = "root"
)

func main() {
	psqlInfo := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, username, password, databaseName,
	)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatal("Error opening sql: ", err)
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("Error ping: ", err)
	}

	fmt.Println("Connection Database Successful!!")
}
