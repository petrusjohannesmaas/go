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
	MaxRtt   string `json:"max_rtt"`
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

	// Define the query with the WITH clause to find the highest avg_rtt
	query := `
    WITH extracted AS (
        
	SELECT
            id,
            elem->>'identity' AS identity,
            elem->>'max_rtt' AS max_rtt,
            (elem->>'avg_rtt')::numeric AS avg_rtt
        FROM public.report,
        jsonb_array_elements(data) AS elem
        WHERE id = 1
    )
    SELECT
        identity,
        max_rtt
    FROM extracted
    WHERE avg_rtt = (
        SELECT MAX(avg_rtt)
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
		err = rows.Scan(&result.Identity, &result.MaxRtt)
		if err != nil {
			log.Fatalf("Error scanning the row: %v", err)
		}
		fmt.Printf("Identity with the highest avg_rtt: %s, Max Average RTT: %s\n", result.Identity, result.MaxRtt)
	} else {
		fmt.Println("No data found with the highest avg_rtt")
	}
}
