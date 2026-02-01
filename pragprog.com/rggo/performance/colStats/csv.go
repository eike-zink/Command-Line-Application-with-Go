package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"strconv"
)

// statsFunc defines a generic statictical function
type statsFunc func(data []float64) float64

func sum(data []float64) float64 {
	sum := 0.0

	for _, d := range data {
		sum += d
	}
	return sum
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

func csvColumnToSlice(r io.Reader, column int) ([]float64, error) {
	// Create the CSV Reader used to read from CSV files
	cr := csv.NewReader(r)
	// Adjusting for 0 based index of column
	column--
	// Read in all CSV data
	rows, err := cr.ReadAll()
	if err != nil {
		return nil, fmt.Errorf("Cannot read data from file: %w", err)
	}
	var data []float64

	// Looping through all records
	// Skip first line
	for i, row := range rows {
		if i == 0 {
			continue
		}
		// Checking number of columns in CSV file
		if len(row) <= column {
			// Row does not have that many columns
			return nil, fmt.Errorf("%w: File has only %d columns.", ErrInvalidColumn, len(row))
		}
		// Try to convert data into a float number (value)
		v, err := strconv.ParseFloat(row[column], 64)
		if err != nil {
			return nil, fmt.Errorf("%w: %s", ErrNotNumber, err)
		}
		data = append(data, v)
	}
	// Return the slice of float64 and nil error
	return data, nil
}
