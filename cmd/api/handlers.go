package main

import (
	"log"
	"net/http"
)

type HomeResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Version string `json:"version"`
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	response := HomeResponse{
		Status:  "200",
		Message: "Anzeigen",
		Version: "1.0.0",
	}

	app.writeJSON(w, http.StatusOK, response)
}

func (app *application) Listings(w http.ResponseWriter, r *http.Request) {
	listings, err := app.DB.AllListings()
	if err != nil {
		app.errorJSON(w, err)
		return
	}
	app.writeJSON(w, http.StatusOK, listings)
}

func (app *application) authenticate(w http.ResponseWriter, r *http.Request) {
	// read json payload

	//validate user against db

	// check password

	// create a jwt user
	u := jwtUser{
		ID:        1,
		FirstName: "Admin",
		LastName:  "User",
	}

	tokens, err := app.auth.GenerateTokenPair(&u)
	if err != nil {
		app.errorJSON(w, err)
		return
	}

	log.Println(tokens.Token)

	w.Write([]byte(tokens.Token))
}

// response := []models.Listing{
// 	{
// 		ID:          "98765527625",
// 		PicUrl:      "https://...",
// 		Category:    "Kleidung",
// 		Title:       "Nice Jacket",
// 		Description: "mint condition",
// 		Price:       40_00,
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 	},
// 	{
// 		ID:          "1274563454",
// 		PicUrl:      "https://2...",
// 		Description: "Broken tooth for sale",
// 		Category:    "Körperteile",
// 		Title:       "Broken tooth",
// 		Price:       5_00,
// 		CreatedAt:   time.Now(),
// 		UpdatedAt:   time.Now(),
// 	},
// }
