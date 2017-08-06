// Package db is a wrapper around the mgo package the app uses to
// communicate with the MongoDB database.
package db

import (
	"log"

	"gopkg.in/mgo.v2"
)

// session is a global variable that persists for the duration of the open
// connection. The app should only need one open connection at a time.
var session *mgo.Session

// InitDB will open the app's connection to the database.
func InitDB() {
	var err error
	session, err = mgo.Dial("database")
	if err != nil {
		panic(err)
	}
}

// CloseDB closes the app's connection with the database.
func CloseDB() {
	log.Println("Closing DB connection")
	session.Close()
}
