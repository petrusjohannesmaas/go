package main

import (
	"testing"
)

// TestConnectPostgresDB tests the database connection
func TestConnectPostgresDB(t *testing.T) {
	db := connectPostgresDB()
	defer db.Close()

	if err := db.Ping(); err != nil {
		t.Errorf("Database connection failed: %v", err)
	}
}

// TestInsert tests the Insert function by verifying data insertion into the database
func TestInsert(t *testing.T) {
	db := connectPostgresDB()
	defer db.Close()

	// Test data
	id := 1
	name := "Test User"
	email := "test@example.com"
	faction := "Test Faction"

	// Insert the data using the Insert function
	insertIntoPostgres(db, id, name, email, faction)

	// Verify the insertion
	var resultName, resultEmail, resultFaction string
	err := db.QueryRow("SELECT name, email, faction FROM test.users WHERE id=$1", id).Scan(&resultName, &resultEmail, &resultFaction)
	if err != nil {
		t.Fatalf("Failed to query inserted data: %v", err)
	}

	if resultName != name || resultEmail != email || resultFaction != faction {
		t.Errorf("Inserted data mismatch: got (%s, %s, %s), expected (%s, %s, %s)", resultName, resultEmail, resultFaction, name, email, faction)
	}

	// Cleanup
	db.Exec("DELETE FROM test.users WHERE id=$1", id)
}
