package main

import (
	"log"
	"net/http"

	"gses2.app-btc/api/handler"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Define the routes
	r.HandleFunc("/api/rate", handler.GetRateHandler).Methods("GET")
	r.HandleFunc("/api/subscribe", handler.SubscribeHandler).Methods("POST")
	r.HandleFunc("/api/sendEmails", handler.SendMailsHandler).Methods("POST")

	port := ":8080"

	server := &http.Server{
		Addr:    port,
		Handler: r,
	}
	// Start the server
	log.Println("Server started on port ", port)
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal("Server error: ", err)
	}

}
