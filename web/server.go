package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nojnhuh/dotbook/db"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "%s", db.GetDotbook("Colts 2015 1-13"))
}

func InitServer(port int) {
	http.HandleFunc("/", handler)
	log.Printf("Ready to serve HTTP on port %d", port)
	path := fmt.Sprintf(":%d", port)
	http.ListenAndServe(path, nil)
}
