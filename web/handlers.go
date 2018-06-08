package web

// var dotbookHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "application/json")
// 	switch r.Method {
// 	case "GET":
// 		vars := mux.Vars(r)
// 		book := db.GetDotbook(vars["id"])
// 		var dotDetails []*models.DotDetails
// 		for _, d := range book.Dots {
// 			dotDetails = append(dotDetails, d.GetDetails(book.Field))
// 		}
// 		data := struct {
// 			Name string               `json:"name"`
// 			Dots []*models.DotDetails `json:"dots"`
// 		}{book.Name, dotDetails}
// 		json.NewEncoder(w).Encode(data)
// 	case "DELETE":
// 		vars := mux.Vars(r)
// 		err := db.DeleteDotbook(vars["id"])
// 		if err != nil {
// 			panic(err)
// 		}
// 	default:
// 		// "PUT" or "POST"
// 		t := struct {
// 			Name  string        `json:"name"`
// 			Field string        `json:"field"`
// 			Dots  []*models.Dot `json:"dots"`
// 		}{}
// 		err := json.NewDecoder(r.Body).Decode(&t)
// 		if err != nil {
// 			panic(err)
// 		}
// 		var field *models.FieldLayout
// 		switch t.Field {
// 		case "HS":
// 			field = models.MakeHSFootball(8, 5)
// 		case "NCAA":
// 			field = models.MakeNCAAFootball(8, 5)
// 		default:
// 			panic("Unsupported field type")
// 		}
// 		book, err := models.NewDotbook(t.Name, field)
// 		if err != nil {
// 			panic(err)
// 		}
// 		if r.Method == "PUT" {
// 			if db.DotbookExists(t.Name) {
// 				err = db.UpdateDotbook(mux.Vars(r)["id"], book)
// 			} else {
// 				err = db.CreateDotbook(book)
// 			}
// 			if err != nil {
// 				panic(err)
// 			}
// 		} else if r.Method == "POST" {
// 			db.CreateDotbook(book)
// 		}
// 		resp := struct {
// 			Message string `json:"message"`
// 		}{"Success", book.ID}
// 		json.NewEncoder(w).Encode(resp)
// 	}

// })

// var dotHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	book := db.GetDotbook(vars["id"])
// 	t := struct {
// 		Name       string  `json:"name"`
// 		MoveCounts float64 `json:"moveCounts"`
// 		HoldCounts float64 `json:"holdCounts"`
// 		XDot       string  `json:"xDot"`
// 		YDot       string  `json:"yDot"`
// 		BodyCenter bool    `json:"bodyCenter"`
// 	}{}
// 	err := json.NewDecoder(r.Body).Decode(&t)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = book.AddDot(t.Name, t.MoveCounts, t.HoldCounts, t.XDot, t.YDot, t.BodyCenter)
// 	if err != nil {
// 		panic(err)
// 	}
// 	db.CreateDotbook(book)
// 	resp := struct {
// 		Message string `json:"message"`
// 		ID      string `json:"_id"`
// 	}{"Success", ""}
// 	json.NewEncoder(w).Encode(resp)
// })

// var indexHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-type", "application/json")
// 	books := db.GetAllDotbooks()
// 	json.NewEncoder(w).Encode(books)
// })

// var tokenHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 	token := jwt.New(jwt.SigningMethodHS256)

// 	// claims := token.Claims.(jwt.MapClaims)
// 	// claims["admin"] = true

// 	tokenString, _ := token.SignedString(mySigningKey)

// 	w.Write([]byte(tokenString))
// })
