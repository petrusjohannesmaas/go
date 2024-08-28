package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

type Product struct {
	Name      string
	Price     float64
	Available bool
}

func main() {
	// Connection string for PostgreSQL
	connStr := "postgres://postgres:secret@localhost:5432/pgtest?sslmode=disable"

	// Open the database connection
	db, err := sql.Open("postgres", connStr) // Use the "postgres" driver
	if err != nil {
		log.Fatal(err)
	}

	// Check the database connection
	if err = db.Ping(); err != nil {
		// Use a type assertion to handle specific SQL errors
		if err == sql.ErrConnDone {
			log.Fatalf("Database connection is already closed or unusable: %v", err)
		} else {
			log.Fatalf("Error pinging the database: %v", err)
		}
	}

	fmt.Println("Successfully connected to the database!")

	createProductTable(db)

	product := Product{"iPhone", 999.99, true}
	pk := insertProduct(db, product)

	fmt.Println("Product inserted with ID:", pk)

	data := []Product{}
	rows, err := db.Query("SELECT name, price, available FROM product")
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()

	var (
		name      string
		price     float64
		available bool
	)

	for rows.Next() {
		err := rows.Scan(&name, &price, &available)
		if err != nil {
			log.Fatal(err)
		}
		data = append(data, Product{name, price, available})
	}

	fmt.Println(data)

}

func createProductTable(db *sql.DB) {

	query := `CREATE TABLE IF NOT EXISTS product (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    price NUMERIC(6,2) NOT NULL,
    available BOOLEAN,
    created timestamp DEFAULT NOW()
	)`

	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
}

func insertProduct(db *sql.DB, product Product) int {
	query := `INSERT INTO product (name, price, available)
		VALUES ($1, $2, $3) RETURNING id`

	var pk int
	err := db.QueryRow(query, product.Name, product.Price, product.Available).Scan(&pk)

	if err != nil {
		log.Fatal(err)
	}
	return pk
}
