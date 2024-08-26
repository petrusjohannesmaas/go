package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Ensure the pq driver is imported
)

var (
	id      int
	name    string
	email   string
	faction string
)

// connecting to the database
func connectPostgresDB() *sql.DB {
	connstring := "user=postgres dbname=test password='blouroomys123' host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		log.Fatal(err)
	}
	return db
}

// Insert takes user inputs and calls insertIntoPostgres
func Insert(db *sql.DB) {
	fmt.Print("Enter user ID: ")
	fmt.Scan(&id)

	fmt.Print("Enter user name: ")
	fmt.Scan(&name)

	fmt.Print("Enter user email: ")
	fmt.Scan(&email)

	fmt.Print("Enter user faction: ")
	fmt.Scan(&faction)

	insertIntoPostgres(db, id, name, email, faction)
}

// insertIntoPostgres inserts user data into the Postgres database
func insertIntoPostgres(db *sql.DB, id int, name string, email string, faction string) {
	_, err := db.Exec("INSERT INTO test.users(id, name, email, faction) VALUES($1, $2, $3, $4)", id, name, email, faction)
	if err != nil {
		fmt.Println("Error inserting data:", err)
	} else {
		fmt.Println("A new user has been created")
	}
}

func main() {
	db := connectPostgresDB()
	defer db.Close()

	Insert(db)
}
