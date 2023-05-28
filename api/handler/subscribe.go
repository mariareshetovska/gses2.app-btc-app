package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os"

	"gses2.app-btc/services"
)

func SubscribeHandler(w http.ResponseWriter, r *http.Request) {
	var requestData services.Record

	err := json.NewDecoder(r.Body).Decode(&requestData)
	if err != nil {
		http.Error(w, "Invalid request", http.StatusBadRequest)
		return
	}

	defer r.Body.Close()

	email := requestData.Email
	if email == "" {
		http.Error(w, "Email is required", http.StatusBadRequest)
		return
	}

	filename := "database.csv"

	err = services.SubscribeEmail(email, filename)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			http.Error(w, "Internal server error", http.StatusInternalServerError)
		} else if errors.Is(err, services.ErrorEmailExist) {
			http.Error(w, "Email already exists", http.StatusConflict)
		} else {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		return
	}

	fmt.Fprintln(w, "Email is added")
}
