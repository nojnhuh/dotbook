// The web package conatins the web server
package web

import (
	"fmt"
	"log"
	"net/http"

	"github.com/nojnhuh/dotbook/db"
)

// handler simply retrieves a default dotbook and displays it.
func handler(w http.ResponseWriter, r *http.Request) {
	book := db.GetDotbook("Colts 2015 1-13")
	fmt.Fprintf(w, "%s", book)
}

// InitServer starts the web serve and declares handler functions
func InitServer(port int) {
	http.HandleFunc("/", handler)
	log.Printf("Ready to serve HTTP on port %d", port)
	path := fmt.Sprintf(":%d", port)
	http.ListenAndServe(path, nil)
}
