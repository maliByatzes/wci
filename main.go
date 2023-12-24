package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func count(r io.Reader, countLines bool) int {
	// A scanner is used to read text from a Reader (such as files)
	scanner := bufio.NewScanner(r)

	// If the count lines flag is not set, we want to count words so we define
	// the scanner split type to words, the default being split by lines
	if !countLines {
		scanner.Split(bufio.ScanWords)
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

	// Parsing the flags provided by the user
	flag.Parse()

	// Calling the count func to count the number of words
	// received from the Standard input and printing it out
	fmt.Println(count(os.Stdin, *lines))
}
