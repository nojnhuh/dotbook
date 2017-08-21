package db

import (
	"github.com/nojnhuh/dotbook/models"
	"gopkg.in/mgo.v2/bson"
)

// DotbookExists returns true if a dotbook with the same name already exists in
// the collection
func DotbookExists(name string) bool {
	c := session.DB("dotbook").C("dotbooks")
	n, err := c.Find(bson.M{"name": name}).Count()
	if err != nil {
		panic(err)
	}
	return n >= 1
}

func GetDotbookId(name string) string {
	c := session.DB("dotbook").C("dotbooks")
	var db *models.Dotbook
	err := c.Find(bson.M{"name": name}).One(&db)
	if err != nil {
		panic(err)
	}
	return ""
}
