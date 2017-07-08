package db

import (
	"log"

	"github.com/nojnhuh/dotbook/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GetDotbook retrieves a dotbook from the database by name and returns a
// pointer to it.
func GetDotbook(name string) *models.Dotbook {
	db := models.Dotbook{}
	c := session.DB("test").C("dotbooks")
	err := c.Find(bson.M{"name": name}).One(&db)
	if err != nil {
		if err == mgo.ErrNotFound {
			log.Println(name, "not found. Try again.")
			return nil
		}
		log.Fatal(err)
	}

	return &db
}
