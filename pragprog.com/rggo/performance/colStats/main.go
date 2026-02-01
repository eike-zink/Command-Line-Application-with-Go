package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Verify and parse arguments
	op := flag.String("op", "sum", "Operation to be executed")
	col := flag.Int("col", 1, "CSV Column on which the Operation executed")

	flag.Parse()

	if err := run(flag.Args(), *op, *col, os.Stdout); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func run(filenames []string, op string, col int, out io.Writer) error {
	var operation statsFunc

	if len(filenames) == 0 {
		return ErrNoFiles
	}

	if col < 1 {
		return fmt.Errorf("%w: %d", ErrInvalidColumn, col)
	}

	switch op {
	case "sum":
		operation = sum
	case "avg":
		operation = avg
	default:
		return fmt.Errorf("%w: %s", ErrInvalidOperation, op)
	}

	consolidate := make([]float64, 0)

	// Loop through all files adding their data to consolidate
	for _, filename := range filenames {
		// Open then file for reading
		file, err := os.Open(filename)
		if err != nil {
			return fmt.Errorf("Cannot open file: %w", err)
		}
		// Parse the CSV into a slice of float64 numbers
		data, err := csvColumnToSlice(file, col)
		if err != nil {
			return err
		}
		if err := file.Close(); err != nil {
			return err
		}

		// Append the data to consolidate
		consolidate = append(consolidate, data...)
	}
	_, err := fmt.Fprintln(out, operation(consolidate))
	return err
}
