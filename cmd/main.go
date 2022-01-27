package main

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Hello World!")
	checkDatabaseConnection()
}

func checkDatabaseConnection() {
	return
	// connection string
	connStr := "user=postgres dbname=connect-db password=secure-password host=localhost sslmode=disable"
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
