package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Data struct {
	Reading string
}

func Extractor() {
	// Connect and open PostgreSQL
	connStr := "postgres://postgres:test123@localhost:5432/pool?sslmode=disable"

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}

	defer db.Close()

	// Check the database connection
	if err = db.Ping(); err != nil {
		if err == sql.ErrConnDone {
			log.Fatalf("Database connection is already closed or unusable: %v", err)
		} else {
			log.Fatalf("Error pinging the database: %v", err)
		}
	}
	fmt.Println("Successfully connected to the database!")

	// Here we run a select statement from within the extractor function
	data := []Data{}
	rows, err := db.Query("SELECT reading FROM data")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	fmt.Println(data)
}
