package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"text/template"

	_ "github.com/mattn/go-sqlite3"
)

type JournalEntry struct {
	ID           int
	BollingerBand string
	HeikenAshi    string
	CandleType    string
	LotSize       int
	Risk          int
}

var tmpl = template.Must(template.ParseFiles("main.html"))

func initDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite3", "./journal.db")
	if err != nil {
		return nil, err
	}

	query := `
	CREATE TABLE IF NOT EXISTS journal (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		bollinger_band TEXT,
		heiken_ashi TEXT,
		candle_type TEXT,
		lot_size INTEGER,
		risk INTEGER
	);
	`
	_, err = db.Exec(query)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := initDB()
	if err != nil {
		fmt.Println("Error initializing database:", err)
		return
	}
	defer db.Close()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			bollingerBand := r.FormValue("bollinger_band")
			heikenAshi := r.FormValue("heiken_ashi")
			candleType := r.FormValue("candle_type")
			lotSize := r.FormValue("lot_size")
			risk := r.FormValue("risk")

			_, err := db.Exec("INSERT INTO journal (bollinger_band, heiken_ashi, candle_type, lot_size, risk) VALUES (?, ?, ?, ?, ?)",
				bollingerBand, heikenAshi, candleType, lotSize, risk)
			if err != nil {
				http.Error(w, "Error inserting data", http.StatusInternalServerError)
				return
			}
		}

		rows, err := db.Query("SELECT id, bollinger_band, heiken_ashi, candle_type, lot_size, risk FROM journal")
		if err != nil {
			http.Error(w, "Error retrieving data", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		var entries []JournalEntry
		for rows.Next() {
			var entry JournalEntry
			err := rows.Scan(&entry.ID, &entry.BollingerBand, &entry.HeikenAshi, &entry.CandleType, &entry.LotSize, &entry.Risk)
			if err != nil {
				http.Error(w, "Error scanning data", http.StatusInternalServerError)
				return
			}
			entries = append(entries, entry)
		}

		tmpl.Execute(w, entries)
	})

	fmt.Println("Server is running on port 8080")
	http.ListenAndServe(":8080", nil)
}

