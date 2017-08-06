package main

import (
	"log"
	"os"

	"github.com/nojnhuh/dotbook/db"
	"github.com/nojnhuh/dotbook/models"
	"github.com/nojnhuh/dotbook/web"
)

func main() {
	dotbook, err := models.NewDotbook("Colts 2015 1-13", models.MakeNCAAFootball(8, 5))
	if err != nil {
		os.Exit(1)
	}
	dotbook.AddDot("1", 0, 8, "1.5 I A40", "2.5 B FSL", false)
	dotbook.AddDot("2", 32, 0, "3.75 O A35", "13 F FH", false)
	dotbook.AddDot("3", 6, 50, "3.75 O A35", "3.5 F FH", false)
	dotbook.AddDot("4", 16, 0, "2.75 I A20", "13.25 F FH", false)
	dotbook.AddDot("4A", 8, 0, "0.25 I A20", "7.5 B FSL", false)
	dotbook.AddDot("5", 8, 0, "4 O A30", "6.5 B FSL", false)
	dotbook.AddDot("5A", 8, 0, "0.75 O A25", "15.25 B FSL", true)
	dotbook.AddDot("5B", 6, 0, "2.5 O A30", "14.5 F FH", true)
	dotbook.AddDot("6", 2, 0, "0 O A30", "14.25 F FH", true)
	dotbook.AddDot("7", 8, 0, "2.25 O A35", "11.75 F FH", false)
	dotbook.AddDot("8", 8, 32, "0.25 O A35", "14.5 F FH", false)
	dotbook.AddDot("9", 8, 0, "3.25 I A30", "12.75 F FH", true)
	dotbook.AddDot("9A", 4, 0, "0.25 I A35", "11.75 F FH", false)
	dotbook.AddDot("10", 4, 0, "3.5 O A40", "11.25 F FH", true)
	dotbook.AddDot("11", 8, 0, "2.5 O A45", "9 F FH", false)
	dotbook.AddDot("12", 8, 0, "3 O A50", "6 F FH", true)
	dotbook.AddDot("13", 8, 0, "3 O A50", "14.5 B FSL", false)

	dotbook2, err := models.NewDotbook("Colts 2015 1-13 2", models.MakeNCAAFootball(8, 5))
	if err != nil {
		os.Exit(1)
	}
	dotbook2.AddDot("1", 0, 8, "1.5 I A40", "2.5 B FSL", false)
	dotbook2.AddDot("2", 32, 0, "3.75 O A35", "13 F FH", false)
	dotbook2.AddDot("3", 6, 50, "3.75 O A35", "3.5 F FH", false)
	dotbook2.AddDot("4", 16, 0, "2.75 I A20", "13.25 F FH", false)
	dotbook2.AddDot("4A", 8, 0, "0.25 I A20", "7.5 B FSL", false)
	dotbook2.AddDot("5", 8, 0, "4 O A30", "6.5 B FSL", false)
	dotbook2.AddDot("5A", 8, 0, "0.75 O A25", "15.25 B FSL", true)
	dotbook2.AddDot("5B", 6, 0, "2.5 O A30", "14.5 F FH", true)
	dotbook2.AddDot("6", 2, 0, "0 O A30", "14.25 F FH", true)
	dotbook2.AddDot("7", 8, 0, "2.25 O A35", "11.75 F FH", false)
	dotbook2.AddDot("8", 8, 32, "0.25 O A35", "14.5 F FH", false)
	dotbook2.AddDot("9", 8, 0, "3.25 I A30", "12.75 F FH", true)
	dotbook2.AddDot("9A", 4, 0, "0.25 I A35", "11.75 F FH", false)
	dotbook2.AddDot("10", 4, 0, "3.5 O A40", "11.25 F FH", true)
	dotbook2.AddDot("11", 8, 0, "2.5 O A45", "9 F FH", false)
	dotbook2.AddDot("12", 8, 0, "3 O A50", "6 F FH", true)
	dotbook2.AddDot("13", 8, 0, "3 O A50", "14.5 B FSL", false)

	db.InitDB()

	defer db.CloseDB()
	db.PersistDotbook(dotbook)
	db.PersistDotbook(dotbook2)

	log.Printf("Found %v dotbooks", len(db.GetAllDotbooks()))

	web.InitServer(8080)
}
