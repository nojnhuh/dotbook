// Package db is a wrapper around the mgo package the app uses to
// communicate with the MongoDB database.
package db

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

// session is a global variable that persists for the duration of the open
// connection. The app should only need one open connection at a time.
var session *mgo.Session

// InitDB will open the app's connection to the database.
func InitDB() {
	var err error
	dbHostname := os.Getenv("DB_DB_HOST")
	dbPort := os.Getenv("DB_DB_PORT")
	dbPath := fmt.Sprintf("%s:%s", dbHostname, dbPort)
	session, err = mgo.Dial(dbPath)
	if err != nil {
		panic(err)
	}
	log.Println("Connected to database at", dbPath)
}

// CloseDB closes the app's connection with the database.
func CloseDB() {
	log.Println("Closing DB connection")
	session.Close()
}
