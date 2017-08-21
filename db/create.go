package db

import (
	"github.com/nojnhuh/dotbook/models"
)

// CreateDotbook stores a Dotbook object in the database.
func CreateDotbook(db *models.Dotbook) error {
	c := session.DB("dotbook").C("dotbooks")
	return c.Insert(db)
}
