package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := mux.NewRouter()

	fileServer := generateFileServer()
	r.HandleFunc("/api/teams", teamHandler)

	r.PathPrefix("/").Handler(http.StripPrefix("/", fileServer))

	server := generateHTTPServer(r, port)

	log.Printf("Listening on port %s...", port)
	log.Fatal(server.ListenAndServe())
}

func generateHTTPServer(r *mux.Router, port string) *http.Server {
	return &http.Server{
		Handler:      r,
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func generateFileServer() http.Handler {
	var dist string
	flag.StringVar(&dist, "dist", "./public/dist", "the directory to serve files from. Defaults to the current dist")
	flag.Parse()
	return http.FileServer(http.Dir(dist))
}