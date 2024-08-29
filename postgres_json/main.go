package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq"
)

// Structs for handling JSON data
type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Income struct {
	Amount      float64 `json:"amount"`
	Description string  `json:"description"`
}

// Connect to the PostgreSQL database
func connectDB() (*sql.DB, error) {
	connStr := "postgres://postgres:secret@localhost:5432/pgtest?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to the database: %v", err)
	}
	return db, nil
}

// Create the PostgreSQL schema if it doesn't exist
func createSchema(db *sql.DB) error {
	schema := `
    CREATE TABLE IF NOT EXISTS users (
        id SERIAL PRIMARY KEY,
        data JSONB NOT NULL
    );

    CREATE TABLE IF NOT EXISTS income (
        id SERIAL PRIMARY KEY,
        data JSONB NOT NULL,
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

// Handler to serve the index page with income records
func indexHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		tmpl := template.Must(template.ParseFiles("index.html"))
		incomes, err := fetchAllIncome(db)
		if err != nil {
			http.Error(w, "Failed to fetch incomes", http.StatusInternalServerError)
			return
		}
		tmpl.Execute(w, map[string][]Income{"Incomes": incomes})
	}
}

// Handler to add a new income record
func addIncomeHandler(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		amount := r.PostFormValue("amount")
		description := r.PostFormValue("description")

		newIncome := Income{
			Amount:      parseFloat(amount),
			Description: description,
		}

		err := addIncome(db, newIncome)
		if err != nil {
			http.Error(w, "Failed to add income", http.StatusInternalServerError)
			return
		}

		tmpl := template.Must(template.ParseFiles("index.html"))
		tmpl.ExecuteTemplate(w, "income-list-element", newIncome)
	}
}

// Function to fetch all income records from the database
func fetchAllIncome(db *sql.DB) ([]Income, error) {
	query := `SELECT data FROM income;`
	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch income: %v", err)
	}
	defer rows.Close()

	var incomes []Income
	for rows.Next() {
		var data []byte
		if err := rows.Scan(&data); err != nil {
			return nil, fmt.Errorf("failed to scan row: %v", err)
		}

		var income Income
		if err := json.Unmarshal(data, &income); err != nil {
			return nil, fmt.Errorf("failed to unmarshal income: %v", err)
		}
		incomes = append(incomes, income)
	}

	return incomes, nil
}

// Function to add a new income record to the database
func addIncome(db *sql.DB, income Income) error {
	incomeJSON, err := json.Marshal(income)
	if err != nil {
		return fmt.Errorf("failed to marshal income: %v", err)
	}

	query := `INSERT INTO income (data) VALUES ($1);`
	_, err = db.Exec(query, incomeJSON)
	if err != nil {
		return fmt.Errorf("failed to insert income: %v", err)
	}

	fmt.Println("Income record added successfully.")
	return nil
}

// Helper function to parse float values safely
func parseFloat(val string) float64 {
	result, err := strconv.ParseFloat(val, 64)
	if err != nil {
		log.Printf("Failed to parse float: %v", err)
		return 0
	}
	return result
}

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatalf("failed to connect to the database: %v", err)
	}
	defer db.Close()

	if err := createSchema(db); err != nil {
		log.Fatalf("failed to create schema: %v", err)
	}

	http.HandleFunc("/", indexHandler(db))
	http.HandleFunc("/add-income/", addIncomeHandler(db))

	fmt.Println("Server is running on port 80")
	http.ListenAndServe("localhost:8000", nil)
}
