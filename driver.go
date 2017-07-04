package main

import (
	"github.com/nojnhuh/dotbook/db"
	"github.com/nojnhuh/dotbook/web"
)

func main() {
	db.DbInit()
	defer db.DbClose()

	web.InitServer(8080)
}
