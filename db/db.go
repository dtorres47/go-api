package db

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func NewDB() (*sql.DB, error) {
	// force SSL off
	dsn := fmt.Sprintf(
		"host=localhost port=5432 user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_NAME"),
	)
	fmt.Println("Connecting with DSN:", dsn)
	return sql.Open("postgres", dsn)
}
