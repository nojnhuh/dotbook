package db

import (
	"github.com/nojnhuh/dotbook/models"
)

func PersistDotbook(db *models.Dotbook) {
	c := session.DB("test").C("dotbooks")
	err := c.Insert(db)
	if err != nil {
		panic(err)
	}
}
