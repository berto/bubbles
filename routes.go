package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"

	_ "github.com/lib/pq"
)

func teamHandler(w http.ResponseWriter, r *http.Request) {
	db, err := connectDB()
	if err != nil {
		log.Fatalf("Could not connect to DB: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	teams, err := getAll(db)
	if err != nil {
		log.Fatalf("Could not query DB: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	teamJSON, err := json.Marshal(teams)
	if err != nil {
		log.Fatalf("Could not parse JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(teamJSON)
}

func connectDB() (*sql.DB, error) {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://localhost/bubbles?sslmode=disable"
	}
	return sql.Open("postgres", connStr)
}
