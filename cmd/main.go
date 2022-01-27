package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	"os"
)

func main() {
	checkDatabaseConnection()
}

func checkDatabaseConnection() {
	// connection string
	connStr := fmt.Sprintf("user=%s dbname=%s password=%s host=%s sslmode=disable",
		os.Getenv("POSTGRES_DB_USERNAME"),
		os.Getenv("POSTGRES_DB_NAME"),
		os.Getenv("POSTGRES_DB_PASSWORD"),
		os.Getenv("POSTGRES_DB_HOST"),
	)
	// connect to the database
	db, err := sql.Open("postgres", connStr)
	// If error is returned then panic
	if err != nil {
		panic(err)
	}
	// This will run db.close() before the function call ends
	defer db.Close()

	// Check if we are able to ping the database
	err = db.Ping()
	if err != nil {
		panic(err)
	}
	// Display success message
	fmt.Printf("\nSuccessfully connected to database!\n")
}
