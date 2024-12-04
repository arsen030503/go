package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
	connStr := "user=postgres password=1337 dbname=go_crud sslmode=disable"
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Unable to connect to database: %v", err)
	}

	err = DB.Ping()
	if err != nil {
		log.Fatalf("Unable to ping database: %v", err)
	}
	fmt.Println("Database connected!")
}
