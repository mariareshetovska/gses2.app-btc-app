package handler

import (
	"gses2.app-btc/services"

	"fmt"
	"log"
	"net/http"
)

const (
	csvFilename    = "database.csv"
	templatePath   = "./email.html"
	emailSubject   = "Bitcoin Rate"
	successMessage = "Email sent successfully"
)

func SendMailsHandler(w http.ResponseWriter, r *http.Request) {
	// Read email addresses from CSV
	emails, err := services.ReadRecordsFromCSV(csvFilename)
	if err != nil {
		log.Fatalf("Failed to read emails from CSV: %v", err)
	}

	// Extract email addresses from records
	to := make([]string, len(emails))
	for i, record := range emails {
		to[i] = record.Email
	}

	// Send email
	err = services.SendBitcoinRateEmail(emailSubject, templatePath, to)
	if err != nil {
		log.Printf("Failed to send email: %v", err)
		http.Error(w, "Failed to send email", http.StatusInternalServerError)
		return
	}

	log.Println(successMessage)
	fmt.Fprintln(w, successMessage)
}
