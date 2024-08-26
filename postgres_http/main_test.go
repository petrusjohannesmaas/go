package main

import (
	"database/sql"
    _ "github.com/lib/pq"
	"testing"
)

func TestConnectPostgresDB(t *testing.T) {
	// Arrange
	connstring := "user=postgres dbname=postgres password='*****' host=localhost port=5432 sslmode=disable"

	// Act
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		t.Errorf("Error connecting to Postgres database: %v", err)
	}
	defer db.Close()

	// Assert
	// Add your assertions here, e.g., check if the database is connected
	if db == nil {
		t.Error("Database connection failed")
	}
}
