package utils

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Data struct {
	Reading string `json:"reading"`
}

func Extractor() ([]byte, error) {
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
	// rows, err := db.Query("SELECT reading ->> 'age' FROM data")
	rows, err := db.Query("SELECT reading FROM data")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	// Scan rows into Data structs
	for rows.Next() {
		var d Data
		err := rows.Scan(&d.Reading)
		if err != nil {
			return nil, err // Return error if scan fails
		}
		data = append(data, d)

	}

	fmt.Printf("The data is %v", data)

	// Encode data as JSON
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err // Return error if encoding fails
	}

	return jsonData, nil // Return encoded JSON data and nil error

}
