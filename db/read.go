package db

import (
	"log"

	"github.com/nojnhuh/dotbook/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// GetAllDotbooks returns a slice of every dotbook in the database
func GetAllDotbooks() []*models.Dotbook {
	c := session.DB("dotbook").C("dotbooks")
	var dbs []*models.Dotbook
	err := c.Find(nil).All(&dbs)
	if err != nil {
		log.Println("Error retrieving dotbooks.")
		return nil
	}
	c.Find(nil).All(&dbs)
	return dbs
}

// GetDotbook retrieves a dotbook from the database by name and returns a
// pointer to it.
func GetDotbook(id string) *models.Dotbook {
	db := models.Dotbook{}
	c := session.DB("dotbook").C("dotbooks")
	err := c.FindId(bson.ObjectIdHex(id)).One(&db)
	if err != nil {
		if err == mgo.ErrNotFound {
			log.Println(id, "not found.")
			return nil
		}
		log.Fatal(err)
	}

	return &db
}
