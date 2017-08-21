package db

import (
	"github.com/nojnhuh/dotbook/models"
	"gopkg.in/mgo.v2/bson"
)

// UpdateDotbook updates the specified dotbook according to the provided data
func UpdateDotbook(id string, data *models.Dotbook) error {
	c := session.DB("dotbook").C("dotbooks")
	return c.UpdateId(bson.ObjectIdHex(id), data)
}
