// Package db is a wrapper around the mgo package the app uses to
// communicate with the MongoDB database.
package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	// database/sql driver
	_ "github.com/lib/pq"
)

// session is a global variable that persists for the duration of the open
// connection. The app should only need one open connection at a time.
var db *sql.DB

// InitDB will open the app's connection to the database.
func InitDB() {
	var err error
	dbUser := "dotbook"
	dbPassword := "pgpass"
	dbName := dbUser
	dbHostname := os.Getenv("DB_DB_HOST")
	dbPort := os.Getenv("DB_DB_PORT")
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", dbUser, dbPassword, dbName, dbHostname, dbPort)
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		err = db.Ping()
		if err == nil {
			break
		}
		time.Sleep(100 * time.Millisecond)
	}
	log.Printf("Connected to database %s at %s:%s\n", dbName, dbHostname, dbPort)
}

// CloseDB closes the app's connection with the database.
func CloseDB() {
	log.Println("Closing DB connection")
	db.Close()
}
