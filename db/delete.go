package db

import (
	"gopkg.in/mgo.v2/bson"
)

// DeleteDotbook deletes a dotbook fromt he database
func DeleteDotbook(id string) error {
	c := session.DB("dotbook").C("dotbooks")
	return c.RemoveId(bson.ObjectIdHex(id))
}
