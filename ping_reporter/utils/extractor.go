package utils

import (
	"database/sql"
	"fmt"
	"log"
	"strconv"

	_ "github.com/lib/pq"
)

// Reading struct to map the result
type Reading struct {
	ID       int
	Identity string
	AvgRtt   float64
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

	// SQL query to get the lowest avg_rtt_ms for each record
	query := `
    WITH extracted AS (
        SELECT
            id,
            elem->>'identity' AS identity,
            elem->>'avg_rtt_ms' AS avg_rtt_ms,
            ROW_NUMBER() OVER (PARTITION BY id ORDER BY CAST(elem->>'avg_rtt_ms' AS DOUBLE PRECISION) ASC) AS rn
        FROM router_os_api.report,
        jsonb_array_elements(neighbors) AS elem
        WHERE elem->>'avg_rtt_ms' ~ '^[0-9]+(\.[0-9]+)?$' -- Filter to ensure avg_rtt_ms is numeric
    )
    SELECT id, identity, avg_rtt_ms
    FROM extracted
    WHERE rn = 1; -- Select the lowest avg_rtt_ms per record
	`

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	defer rows.Close()

	// Iterate through each result and print the lowest avg_rtt_ms for each record
	for rows.Next() {
		var reading Reading
		var avgRttStr string

		err = rows.Scan(&reading.ID, &reading.Identity, &avgRttStr)
		if err != nil {
			log.Printf("Error scanning the row: %v", err)
			continue
		}

		// Convert avg_rtt_ms to float64
		reading.AvgRtt, err = strconv.ParseFloat(avgRttStr, 64)
		if err != nil {
			log.Printf("Error parsing avg_rtt_ms: %v", err)
			continue
		}

		// Print the result for each record
		fmt.Printf("Record ID: %d, Neighbor with lowest avg_rtt_ms: %s, RTT: %.2f ms\n", reading.ID, reading.Identity, reading.AvgRtt)
	}

	// Check for any errors encountered during iteration
	if err = rows.Err(); err != nil {
		log.Fatalf("Error during row iteration: %v", err)
	}
}
