package main

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	_ "github.com/lib/pq"
)

// Database connection string
const connStr = "user=postgres password=blouroomys123 dbname=postgres sslmode=disable"

// Function to generate random networking data
func generateData() (string, string, float64, float64) {
	sourceIP := fmt.Sprintf("192.168.1.%d", rand.Intn(255))
	destinationIP := fmt.Sprintf("192.168.2.%d", rand.Intn(255))
	latency := rand.Float64() * 100   // Random latency in ms
	bandwidth := rand.Float64() * 100 // Random bandwidth in Mbps
	return sourceIP, destinationIP, latency, bandwidth
}

func insertData(db *sql.DB, sourceIP, destinationIP string, latency, bandwidth float64) error {
	query := `
        INSERT INTO network_metrics (time, source_ip, destination_ip, latency_ms, bandwidth_mbps)
        VALUES ($1, $2, $3, $4, $5)
    `
	_, err := db.Exec(query, time.Now(), sourceIP, destinationIP, latency, bandwidth)
	return err
}

func main() {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Generate and insert data every second
	for {
		sourceIP, destinationIP, latency, bandwidth := generateData()
		err := insertData(db, sourceIP, destinationIP, latency, bandwidth)
		if err != nil {
			log.Println("Error inserting data:", err)
		} else {
			log.Println("Inserted data:", sourceIP, destinationIP, latency, bandwidth)
		}
		time.Sleep(1 * time.Second) // Simulate continuous data generation
	}
}
