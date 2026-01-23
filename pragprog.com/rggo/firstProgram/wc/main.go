package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {
	// Define a boolean flag -l to count line instead of words
	countLines := flag.Bool("l", false, "Count liness")
	flag.Parse()

	// Calling the count function to count the number of words
	// received from the Standard Input and printing it out
	fmt.Println(count(os.Stdin, *countLines))
}

func count(r io.Reader, countLines bool) int {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// Define the scanner split type to words or lines (default)
	if !countLines {
		scanner.Split(bufio.ScanWords)
	}

	// Define a counter
	wc := 0

	// For every word scanned, increment th counter
	for scanner.Scan() {
		wc++
	}

	// Return the counter
	return wc
}
