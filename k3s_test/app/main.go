package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

// Function to connect to the PostgreSQL database
func connectDB() (*sql.DB, error) {
	connStr := os.Getenv("DB_URL")

	if connStr == "" {
		return nil, fmt.Errorf("DB_URL environment variable is not set")
	}

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}
	return db, nil
}

// Function to create the schema (users and income tables)
func createSchema(db *sql.DB) error {
	schema := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        username VARCHAR(255) UNIQUE NOT NULL,
        password VARCHAR(255) NOT NULL
    );

    CREATE TABLE IF NOT EXISTS income (
        id SERIAL PRIMARY KEY,
        amount NUMERIC NOT NULL,
        description TEXT,
        user_id INT REFERENCES users(id) ON DELETE CASCADE
    );
    `
	_, err := db.Exec(schema)
	if err != nil {
		return fmt.Errorf("failed to create schema: %v", err)
	}
	fmt.Println("Schema created successfully.")
	return nil
}

// Function to add a new user
func createUser(db *sql.DB, username, password string) (int, error) {
	query := `INSERT INTO users (username, password) VALUES ($1, $2) RETURNING id;`
	var userID int
	err := db.QueryRow(query, username, password).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("failed to create user: %v", err)
	}
	fmt.Printf("User created with ID: %d\n", userID)
	return userID, nil
}

// Function to add a new income record
func createIncome(db *sql.DB, amount float64, description string, userID int) error {
	query := `INSERT INTO income (amount, description, user_id) VALUES ($1, $2, $3);`
	_, err := db.Exec(query, amount, description, userID)
	if err != nil {
		return fmt.Errorf("failed to create income record: %v", err)
	}
	fmt.Println("Income record created successfully.")
	return nil
}

// Function to fetch all income records for a user
func fetchIncomeByUser(db *sql.DB, userID int) error {
	query := `SELECT id, amount, description FROM income WHERE user_id = $1;`
	rows, err := db.Query(query, userID)
	if err != nil {
		return fmt.Errorf("failed to fetch income records: %v", err)
	}
	defer rows.Close()

	fmt.Printf("Income records for user ID %d:\n", userID)
	for rows.Next() {
		var id int
		var amount float64
		var description string
		err = rows.Scan(&id, &amount, &description)
		if err != nil {
			return fmt.Errorf("failed to scan row: %v", err)
		}
		fmt.Printf("ID: %d, Amount: %.2f, Description: %s\n", id, amount, description)
	}

	return nil
}

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	// Create the schema
	if err := createSchema(db); err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}

	// Add two users
	user1ID, err := createUser(db, "user1", "password1")
	if err != nil {
		log.Fatalf("failed to create user1: %v", err)
	}

	user2ID, err := createUser(db, "user2", "password2")
	if err != nil {
		log.Fatalf("failed to create user2: %v", err)
	}

	// Add income records for each user
	if err := createIncome(db, 5000, "Salary", user1ID); err != nil {
		log.Fatalf("failed to create income for user1: %v", err)
	}

	if err := createIncome(db, 3000, "Freelance", user2ID); err != nil {
		log.Fatalf("failed to create income for user2: %v", err)
	}

	// Fetch and display income records for each user
	if err := fetchIncomeByUser(db, user1ID); err != nil {
		log.Fatalf("failed to fetch income for user1: %v", err)
	}

	if err := fetchIncomeByUser(db, user2ID); err != nil {
		log.Fatalf("failed to fetch income for user2: %v", err)
	}
}
