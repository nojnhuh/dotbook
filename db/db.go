package db

import (
	"log"

	"github.com/nojnhuh/dotbook/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var session *mgo.Session

func DbInit() {
	var err error
	session, err = mgo.Dial("database")
	if err != nil {
		panic(err)
	}
}

func DbClose() {
	log.Println("Closing DB connection")
	session.Close()
}

func PersistDotbook(db *models.Dotbook) {
	c := session.DB("test").C("dotbooks")
	err := c.Insert(db)
	if err != nil {
		log.Fatal(err)
	}
}

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
