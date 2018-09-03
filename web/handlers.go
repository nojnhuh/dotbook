package web

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/nojnhuh/dotbook/db"
	"github.com/nojnhuh/dotbook/models"
)

var dotbookHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	id64, err := strconv.ParseInt(mux.Vars(r)["id"], 10, 0)
	if err != nil {
		log.Fatal(err)
	}
	id := int(id64)
	switch r.Method {
	case "GET":
		book := db.GetDotbook(id)
		if book == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(http.StatusText(http.StatusNotFound)))
			return
		}
		var dotDetails []*models.DotDetails
		for _, d := range book.Dots {
			dotDetails = append(dotDetails, d.GetDetails(book.Field))
		}
		data := struct {
			Name string               `json:"name"`
			Dots []*models.DotDetails `json:"dots"`
		}{book.Name, dotDetails}
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(data)
	case "DELETE":
		err := db.DeleteDotbook(id)
		if err != nil {
			panic(err)
		}
	default:
		// "PUT" or "POST"
		t := struct {
			Name  string        `json:"name"`
			Field string        `json:"field"`
			Dots  []*models.Dot `json:"dots"`
		}{}
		err := json.NewDecoder(r.Body).Decode(&t)
		if err != nil {
			panic(err)
		}
		var field *models.FieldLayout
		switch t.Field {
		case "HS":
			field = models.MakeHSFootball(8, 5)
		case "NCAA":
			field = models.MakeNCAAFootball(8, 5)
		default:
			panic("Unsupported field type")
		}
		book, err := models.NewDotbook(t.Name, field)
		if err != nil {
			panic(err)
		}
		if r.Method == "PUT" {
			if db.DotbookExists(t.Name) {
				err = db.UpdateDotbook(id, book)
			} else {
				db.CreateDotbook(book)
			}
			if err != nil {
				panic(err)
			}
		} else if r.Method == "POST" {
			db.CreateDotbook(book)
		}
		resp := struct {
			Message string `json:"message"`
			ID      int    `json:"id"`
		}{"Success", id}
		w.Header().Set("Content-type", "application/json")
		json.NewEncoder(w).Encode(resp)
	}
})

var dotHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	id64, err := strconv.ParseInt(mux.Vars(r)["db_id"], 10, 0)
	if err != nil {
		log.Fatal(err)
	}
	id := int(id64)
	t := struct {
		Name       string  `json:"name"`
		MoveCounts float64 `json:"moveCounts"`
		HoldCounts float64 `json:"holdCounts"`
		XDot       string  `json:"xDot"`
		YDot       string  `json:"yDot"`
		BodyCenter bool    `json:"bodyCenter"`
	}{}
	err = json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		panic(err)
	}
	book := db.GetDotbook(id)
	err = book.AddDot(t.Name, t.MoveCounts, t.HoldCounts, t.XDot, t.YDot, t.BodyCenter)
	if err != nil {
		panic(err)
	}
	db.CreateDotbook(book)
	resp := struct {
		Message string `json:"message"`
		ID      string `json:"_id"`
	}{"Success", ""}
	w.Header().Set("Content-type", "application/json")
	json.NewEncoder(w).Encode(resp)
})

var indexHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	books := db.GetAllDotbooks()
	json.NewEncoder(w).Encode(books)
})

var tokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	token := jwt.New(jwt.SigningMethodHS256)

	// claims := token.Claims.(jwt.MapClaims)
	// claims["admin"] = true

	tokenString, _ := token.SignedString(mySigningKey)

	w.Write([]byte(tokenString))
})
