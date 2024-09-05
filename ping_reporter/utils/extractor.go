package utils

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Reading struct to map the result
type Reading struct {
	Identity string `json:"identity"`
	AvgRtt   string `json:"avg_rtt_ms"`
}

func Extractor() {
	// Connect and open PostgreSQL
	connStr := "postgres://postgres:blouroomys123@localhost:5432/test?sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error opening database: %v", err)
	}
	defer db.Close()

	// Check the database connection
	if err = db.Ping(); err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}
	fmt.Println("Successfully connected to the database!")

	// Define the query with the WITH clause to find the highest avg_rtt_ms
	query := `
    WITH extracted AS (
        
	SELECT
            id,
            elem->>'identity' AS identity,
            elem->>'avg_rtt_ms' AS avg_rtt_ms
        FROM router_os_api.report,
        jsonb_array_elements(neighbors) AS elem
    )
    SELECT
        identity,
        avg_rtt_ms
    FROM extracted
    WHERE avg_rtt_ms = (
        SELECT MAX(avg_rtt_ms)
        FROM extracted
    );
	`

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	defer rows.Close()

	// Define a variable to store the result
	var result Reading

	// Check if there's at least one row and scan the result
	if rows.Next() {
		err = rows.Scan(&result.Identity, &result.AvgRtt)
		if err != nil {
			log.Fatalf("Error scanning the row: %v", err)
		}
		fmt.Printf("Neighbor with lowest avg_rtt_ms: %s, RTT: %s\n", result.Identity, result.AvgRtt)
	} else {
		fmt.Println("No data found")
	}
}
