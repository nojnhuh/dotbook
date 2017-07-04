package db

import (
	"log"

	"github.com/nojnhuh/dotbook/models"
	"gopkg.in/mgo.v2/bson"
)

func GetDotbook(name string) *models.Dotbook {
	log.Println("Retrieving Dotbook:", "Colts 2015 1-13")
	db := models.Dotbook{}
	c := session.DB("test").C("dotbooks")
	err := c.Find(bson.M{"name": name}).One(&db)
	if err != nil {
		log.Fatal(err)
	}

	return &db
}
