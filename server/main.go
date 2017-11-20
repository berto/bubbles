package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/berto/bubbles/server/routes"
	"github.com/gorilla/mux"
	_ "github.com/joho/godotenv/autoload"
	"github.com/rs/cors"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	r := mux.NewRouter()

	fileServer := generateFileServer()
	r.HandleFunc("/api/teams", routes.TeamHandler)

	r.PathPrefix("/").Handler(routes.NotFoundHook{http.StripPrefix("/", fileServer)})

	server := generateHTTPServer(r, port)

	log.Printf("Listening on port %s...", port)
	log.Fatal(server.ListenAndServe())
}

func generateHTTPServer(r *mux.Router, port string) *http.Server {
	return &http.Server{
		Handler:      cors.Default().Handler(r),
		Addr:         ":" + port,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
}

func generateFileServer() http.Handler {
	var public string
	flag.StringVar(&public, "public", "./public", "the directory to serve files from. Defaults to public")
	flag.Parse()
	return http.FileServer(http.Dir(public))
}
