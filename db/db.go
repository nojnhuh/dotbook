package db

import (
	"log"

	"gopkg.in/mgo.v2"
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
