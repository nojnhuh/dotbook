// Package web contains the web server
package web

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

var (
	r            *mux.Router
	mySigningKey = []byte("secret")
)

func setPaths() {
	r = mux.NewRouter()
	r.Handle("/dotbooks", jwtMiddleware.Handler(indexHandler)).Methods("GET")
	r.Handle("/dotbooks/{id}", jwtMiddleware.Handler(dotbookHandler)).Methods("GET", "POST", "PUT", "DELETE")
	r.Handle("/dotbooks/{db_id}/dots/{dot_id}", jwtMiddleware.Handler(dotHandler)).Methods("GET", "POST", "PUT", "DELETE")
	r.Handle("/token", tokenHandler).Methods("GET")
	http.Handle("/", r)
}

// InitServer starts the web serve and declares handler functions
func InitServer(port int) {
	setPaths()
	log.Printf("Ready to serve HTTP on port %d.", port)
	path := fmt.Sprintf(":%d", port)
	log.Fatal(http.ListenAndServe(path, handlers.LoggingHandler(os.Stdout, r)))
}
