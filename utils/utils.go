package utils

import (
	"encoding/csv"
	"log"
	"os"
)

// CustomData represents a custom interface that must be satisfied by any data type
// that wants to be written to CSV using the WriteToCSV function.
type CustomData interface {
	ToCSVRecord() []string
}

// WriteToCSV writes data to a CSV file.
func WriteToCSV(headers []string, data []CustomData, fileName string) error {
	file, err := os.Create(fileName)
	if err != nil {
		log.Fatalln("Failed to create output CSV file", err)
	}
	defer file.Close()

	writer := csv.NewWriter(file)

	// Write the headers to the CSV file
	err = writer.Write(headers)
	if err != nil {
		return err
	}

	for _, row := range data {
		// Get the CSV record for each data type
		record := row.ToCSVRecord()

		// writing a new CSV record
		err := writer.Write(record)
		if err != nil {
			return err
		}
	}
	writer.Flush()

	return writer.Error()
}
