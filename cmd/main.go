package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"github.com/AbhiiGatty/GoUrlShortner/database"
	"github.com/AbhiiGatty/GoUrlShortner/model"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
	"os"
	"github.com/shomali11/util/xhashes"
)

func main() {
	// Initialize logger to print messages to console
	log.SetFormatter(&log.JSONFormatter{})
	//// Check database connection
	intializeDatabaseConnection()
	// This will run db.close() before the main function call ends
	defer func(db *sql.DB) {
		err := db.Close()
		// If error is returned closing connection to the database then just log the error
		if err != nil {
			errorMessage := fmt.Sprintf("Error: Not able to close connection to database - %s", os.Getenv("POSTGRES_DB_NAME"))
			log.Error(errorMessage)
		}else{
			successMessage := fmt.Sprintf("Successfully disconnected to database - %s", os.Getenv("POSTGRES_DB_NAME"))
			log.Info(successMessage)
		}
	}(database.DBCon)
	
	// We will emulate that we are sending many requests with a long url of varying length
	if os.Getenv("ENVIRONMENT") == model.Local{
		// This will only be done for local development to see some values in database to make
		// sure everything is working as expected
		populateUrlMapTable()
	}
}

func intializeDatabaseConnection() {
	// Initialize the connection string
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s port=%s host=%s sslmode=disable",
		os.Getenv("POSTGRES_DB_USERNAME"),
		os.Getenv("POSTGRES_DB_PASSWORD"),
		os.Getenv("POSTGRES_DB_NAME"),
		os.Getenv("POSTGRES_DB_PORT"),
		os.Getenv("POSTGRES_DB_HOST"),
	)

	// connect to the database
	dbConnection, err := sql.Open("postgres", connStr)
	// Set database connection globally
	database.DBCon = dbConnection

	// If error is returned while connecting to the database then log and raise run time error
	if err != nil {
		errorMessage := fmt.Sprintf("Error: Not able to connect to database - %s", os.Getenv("POSTGRES_DB_NAME"))
		log.Error(errorMessage)
		panic(err)
	}

	// Check if we are able to ping the database
	err = database.DBCon.Ping()
	// If error is returned while pinging the database then log and raise run time error
	if err != nil {
		errorMessage := fmt.Sprintf("Error: Not able to ping the database - %s", os.Getenv("POSTGRES_DB_NAME"))
		log.Error(errorMessage)
		panic(err)
	}
	// Display success message
	log.Info("Successfully connected from database -> " + os.Getenv("POSTGRES_DB_NAME"))
}


func populateUrlMapTable(){
	// Get the file handler on the file which contains a lot of URL samples
	file, err := os.Open(os.Getenv("MOCK_URL_FILE_PATH"))
	// Check for errors while opening
	if err != nil {
		log.Error(err)
		return
	}else{
		// Close file handler before returning function call
		defer file.Close()
	}

	// Read the file content
	scanner := bufio.NewScanner(file)

	// We only want to read N (limit) number of lines
	count := 0
	limit := 10
	// Get the file line by line
	for scanner.Scan() {
		// Create short code for each url
		generateShortUrlCode(scanner.Text())
		// Increment the count and check if it has reached the limit
		count++
		if count > limit {
			break
		}
	}
	// If there are any errors while reading the file then log it and return
	if err := scanner.Err(); err != nil {
		log.Error(err)
	}
}

func generateShortUrlCode(url string) string {
	// Create a SHA256 hash of the url and return it
	// Here we only use the first 7 characters out of 64 character
	shortUrlCode := string(xhashes.SHA256(url)[0:7])
	// Insert the short code into the database
	sqlStatement := `INSERT INTO url_map("fullUrl", "shortUrlCode") VALUES ($1, $2)`
	_, err := database.DBCon.Exec(sqlStatement, url, shortUrlCode)
	if err != nil {
		errorMessage := fmt.Sprintf("Error While inseting: %s -> %s", url, shortUrlCode)
		log.Error(errorMessage)
		log.Error(err)
	}else{
		Successmessage := fmt.Sprintf("Inserted ShortUrlCode: %s for URL: %s", shortUrlCode, url)
		log.Info(Successmessage)
	}
	//Return the short code generate for the URL
	return shortUrlCode
}