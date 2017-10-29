package main

import (
	"database/sql"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"

	"github.com/berto/bubbles/db/queries"
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

	teams, err := queries.GetTeams(db)
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

type hookedResponseWriter struct {
	http.ResponseWriter
	ignore bool
}

func (hrw *hookedResponseWriter) WriteHeader(status int) {
	if status == 404 {
		hrw.ResponseWriter.Header().Set("Content-Type", "text/html")
		index, err := ioutil.ReadFile("./public/dist/index.html")
		hrw.ignore = true
		if err != nil {
			log.Fatal(err)
		}
		hrw.ResponseWriter.WriteHeader(200)
		hrw.ResponseWriter.Write(index)
	}
}

func (hrw *hookedResponseWriter) Write(p []byte) (int, error) {
	if hrw.ignore {
		return len(p), nil
	}
	return hrw.ResponseWriter.Write(p)
}

type NotFoundHook struct {
	h http.Handler
}

func (nfh NotFoundHook) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	nfh.h.ServeHTTP(&hookedResponseWriter{ResponseWriter: w}, r)
}
