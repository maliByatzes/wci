package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countLines bool, countBytes bool) int {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// If the count lines flag or count bytes is not set, we want to count words
	// so we define the scanner split type to words, the default being split by lines
	if !countLines && !countBytes {
		scanner.Split(bufio.ScanWords)
	}

	if countBytes {
		scanner.Split(bufio.ScanBytes)
	}

	// Define the counter
	wc := 0

	// For every word counted increment the counter
	for scanner.Scan() {
		wc++
	}

	// Return the total
	return wc
}

func main() {
	// Define a boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")

	// Define a boolean flag -b to count bytes instad of words or lines
	bytes := flag.Bool("b", false, "Count bytes")

	// Parsing the flags provided by the user
	flag.Parse()

	// Calling the count func to count the number of words
	// received from the Standard input and printing it out
	fmt.Println(count(os.Stdin, *lines, *bytes))
}
