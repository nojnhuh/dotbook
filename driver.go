package main

import (
	"fmt"

	database "github.com/nojnhuh/dotbook/db"
	"github.com/nojnhuh/dotbook/models"
)

var db *models.Dotbook

func main() {
	db = models.NewDotbook("NEW Colts 2015 1-13", models.MakeNCAAFootball(8))
	db.AddDot("1", 0, 8, "1.5 I 1 40", "2.5 B FSL", false)
	db.AddDot("2", 32, 0, "3.75 O 1 35", "13 F FH", false)
	db.AddDot("3", 6, 50, "3.75 O 1 35", "3.5 F FH", false)
	db.AddDot("4", 16, 0, "2.75 I 1 20", "13.25 F FH", false)
	db.AddDot("4A", 8, 0, "0.25 I 1 20", "7.5 B FSL", false)
	db.AddDot("5", 8, 0, "4 O 1 30", "6.5 B FSL", false)
	db.AddDot("5A", 8, 0, "0.75 O 1 25", "15.25 B FSL", true)
	db.AddDot("5B", 6, 0, "2.5 O 1 30", "14.5 F FH", true)
	db.AddDot("6", 2, 0, "0 O 1 30", "14.25 F FH", true)
	db.AddDot("7", 8, 0, "2.25 O 1 35", "11.75 F FH", false)
	db.AddDot("8", 8, 32, "0.25 O 1 35", "14.5 F FH", false)
	db.AddDot("9", 8, 0, "3.25 I 1 30", "12.75 F FH", true)
	db.AddDot("9A", 4, 0, "0.25 I 1 35", "11.75 F FH", false)
	db.AddDot("10", 4, 0, "3.5 O 1 40", "11.25 F FH", true)
	db.AddDot("11", 8, 0, "2.5 O 1 45", "9 F FH", false)
	db.AddDot("12", 8, 0, "3 O 1 50", "6 F FH", true)
	db.AddDot("13", 8, 0, "3 O 1 50", "14.5 B FSL", false)

	// fmt.Println(db)
	database.DbInit()
	defer database.DbClose()
	database.PersistDotbook(db)
	db2 := database.GetDotbook()
	fmt.Println(db2)
}
