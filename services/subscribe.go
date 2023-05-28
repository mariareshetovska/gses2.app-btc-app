package services

import (
	"encoding/csv"
	"errors"
	"os"
)

var ErrorEmailExist = errors.New("email already exists in the database")

type Record struct {
	Email string `json:"email"`
}

func AppendRecordToCSV(record Record, filename string) error {
	// Create or open the CSV file
	file, err := os.OpenFile(filename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)

	// Write the record to the CSV file
	err = writer.Write([]string{record.Email})
	if err != nil {
		return err
	}

	// Flush any buffered data to the underlying writer (the file)
	writer.Flush()

	// Check for any error during the flushing process
	if err := writer.Error(); err != nil {
		return err
	}

	return nil
}

func ReadRecordsFromCSV(filename string) ([]Record, error) {
	// Open the CSV file
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	// Create a slice to hold the parsed records
	result := make([]Record, len(records))
	for i, row := range records {
		if len(row) > 0 {
			result[i] = Record{Email: row[0]}
		}
	}

	return result, nil
}

func SubscribeEmail(email string, filename string) error {
	records, err := ReadRecordsFromCSV(filename)
	if err != nil {
		return err
	}

	for _, record := range records {
		if record.Email == email {
			// Email already exists in the database
			return ErrorEmailExist
		}
	}

	err = AppendRecordToCSV(Record{Email: email}, filename)
	if err != nil {
		return err
	}

	return nil
}
