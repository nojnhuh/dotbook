// Package db is a wrapper around the mgo package the app uses to
// communicate with the MongoDB database.
package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/lib/pq" // database/sql driver
	"github.com/nojnhuh/dotbook/models"
)

// session is a global variable that persists for the duration of the open
// connection. The app should only need one open connection at a time.
var db *sql.DB

// InitDB will open the app's connection to the database.
func InitDB() {
	dbUser := os.Getenv("DB_DB_USER")
	dbPassword := os.Getenv("DB_DB_PASSWORD")
	dbHostname := os.Getenv("DB_DB_HOST")
	dbName := dbUser
	connStr := fmt.Sprintf("user=%s password=%s dbname=%s host=%s sslmode=disable", dbUser, dbPassword, dbName, dbHostname)

	var err error
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
	log.Printf("Connected to database %s at %s\n", dbName, dbHostname)

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS dotbooks (
		id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
		name text
	);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS dots (
		id int GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
		name text
	);`)
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(`CREATE TABLE IF NOT EXISTS dotbooks_dots (
		dotbook_id int not null,
		dot_id int not null,
		primary key(dotbook_id, dot_id)
	);`)
	if err != nil {
		log.Fatal(err)
	}
}

// CloseDB closes the app's connection with the database.
func CloseDB() {
	log.Println("Closing DB connection")
	db.Close()
}

// GetAllDotbooks queries the database and returns a Dotbook slice
func GetAllDotbooks() []*models.Dotbook {
	rows, err := db.Query(`SELECT * FROM dotbooks`)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	dbs := []*models.Dotbook{}

	for rows.Next() {
		var id int
		var name string
		err = rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}

		dotbook := &models.Dotbook{Name: name}
		dbs = append(dbs, dotbook)
	}
	err = rows.Err() // get any error encountered during iteration
	if err != nil {
		log.Fatal(err)
	}

	return dbs
}

// CreateDotbook inserts a Dotbook object in the database
func CreateDotbook(dotbook *models.Dotbook) error {
	newDotIDs := []int{}
	for _, dot := range dotbook.Dots {
		newDotIDs = append(newDotIDs, CreateDot(dot))
	}
	var newDotbookID int
	err := db.QueryRow(`INSERT INTO dotbooks (name) VALUES ($1) RETURNING id`, dotbook.Name).Scan(&newDotbookID)

	for _, id := range newDotIDs {
		db.QueryRow(`INSERT INTO dotbooks_dots VALUES ($1, $2)`, newDotbookID, id)
	}

	return err
}

// CreateDot inserts a Dot object in the database
func CreateDot(dotbook *models.Dot) int {
	var newID int
	err := db.QueryRow(`INSERT INTO dots (name) VALUES ($1) RETURNING id`, dotbook.Name).Scan(&newID)
	if err != nil {
		log.Fatal(err)
	}
	return newID
}
