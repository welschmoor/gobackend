package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/welschmoor/gobackend/internal/models"
)

type HomeResponse struct {
	Status  string `json:"status"`
	Message string `json:"message"`
	Version string `json:"version"`
}

func (app *application) Home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := HomeResponse{
		Status:  "200",
		Message: "Anzeigen",
		Version: "1.0.0",
	}

	out, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}

func (app *application) Listings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	response := []models.Listing{
		{
			ID:          "98765527625",
			PicUrl:      "https://...",
			Category:    "Kleidung",
			Title:       "Nice Jacket",
			Description: "mint condition",
			Price:       40_00,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          "1274563454",
			PicUrl:      "https://2...",
			Description: "Broken tooth for sale",
			Category:    "Körperteile",
			Title:       "Broken tooth",
			Price:       5_00,
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	out, err := json.Marshal(response)
	if err != nil {
		fmt.Println(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(out)
}
