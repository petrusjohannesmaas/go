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
	ID          int
	PPPusername string
	IP          string
	Identity    string
	AvgRtt      float64
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
            ppp_username,
			ip,
            elem->>'identity' AS identity,
            elem->>'avg_rtt_ms' AS avg_rtt_ms,
            ROW_NUMBER() OVER (PARTITION BY id ORDER BY CAST(elem->>'avg_rtt_ms' AS DOUBLE PRECISION) ASC) AS rn
        FROM router_os_api.report,
        jsonb_array_elements(neighbors) AS elem
        WHERE elem->>'avg_rtt_ms' ~ '^[0-9]+(\.[0-9]+)?$' -- Filter to ensure avg_rtt_ms is numeric
    )
    SELECT id, ppp_username, ip, identity, avg_rtt_ms
    FROM extracted
    WHERE rn = 1; -- Select the lowest avg_rtt_ms per record
	`

	// Execute the query
	rows, err := db.Query(query)
	if err != nil {
		log.Fatalf("Error executing query: %v", err)
	}
	defer rows.Close()

	// Slice to store the results
	var readings []Reading

	// Iterate through each result and collect the lowest avg_rtt_ms for each record
	for rows.Next() {
		var reading Reading
		var avgRttStr string

		err = rows.Scan(&reading.ID, &reading.PPPusername, &reading.IP, &reading.Identity, &avgRttStr)
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

		// Append the reading to the slice
		readings = append(readings, reading)
	}

	// Check for any errors encountered during iteration
	if err = rows.Err(); err != nil {
		log.Fatalf("Error during row iteration: %v", err)
	}

	// Export the results to an Excel file
	err = WriteToExcel(readings, "output.xlsx")
	if err != nil {
		log.Fatalf("Error exporting to Excel: %v", err)
	}

	fmt.Println("Data successfully exported to output.xlsx")
}
