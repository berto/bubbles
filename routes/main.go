package routes

import (
	"database/sql"
	"os"
)

func connectDB() (*sql.DB, error) {
	connStr := os.Getenv("DATABASE_URL")
	if connStr == "" {
		connStr = "postgres://localhost/bubbles?sslmode=disable"
	}
	return sql.Open("postgres", connStr)
}
