package main

import (
	"database/sql"
	"fmt"
)

var (
    id int
    name string
    email string
    faction string
)

// connecting to the database
func connectPostgresDB() *sql.DB {
	connstring := "user=postgres dbname=test password='blouroomys123' host=localhost port=5432 sslmode=disable"
	db, err := sql.Open("postgres", connstring)
	if err != nil {
		fmt.Println(err)
	}
	return db
}

func Insert(db *sql.DB) {
    id = ${#id}
    name = ${name}
    email = ${email}
    faction = ${faction}
    insertIntoPostgres(db, id, name, email, faction)
}

func insertIntoPostgres(db *sql.DB, id, name, email, faction) {
    _, err := db.Exec("INSERT INTO test.users(id, name, email, faction) VALUES($1,$2,$3,$4)", id, name, email, faction)
    if err != nil {
        fmt.Println(err)
    } else {
        fmt.Println("A new user has been created")
    }
}
