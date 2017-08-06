package db

import (
	"github.com/nojnhuh/dotbook/models"
)

// PersistDotbook stores a Dotbook object in the database.
func PersistDotbook(db *models.Dotbook) {
	c := session.DB("dotbook").C("dotbooks")
	err := c.Insert(db)
	if err != nil {
		panic(err)
	}
}
