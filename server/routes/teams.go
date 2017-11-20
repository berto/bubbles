package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/berto/bubbles/server/db/queries"
	_ "github.com/lib/pq"
)

func TeamHandler(w http.ResponseWriter, r *http.Request) {
	db, err := queries.ConnectDB()
	if err != nil {
		log.Print("Could not connect to DB: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer db.Close()

	teams, err := queries.GetTeams(db)
	if err != nil {
		log.Print("Could not query DB: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	teamJSON, err := json.Marshal(teams)
	if err != nil {
		log.Print("Could not parse JSON: %v", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(teamJSON)
}
