package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func main() {

	// Defining boolean flag -l to count lines instead of words
	lines := flag.Bool("l", false, "Count lines")

	// Defining boolean flag -b to count bytes instead of words
	bytes := flag.Bool("b", false, "Count bytes")
	flag.Parse()

	wc := count(os.Stdin, *lines, *bytes)
	fmt.Println(wc)
}

func count(r io.Reader, countLines bool, countBytes bool) int {
	scanner := bufio.NewScanner(r)
	if countLines {
		scanner.Split(bufio.ScanLines)
	} else if countBytes {
		scanner.Split(bufio.ScanBytes)
	} else {
		scanner.Split(bufio.ScanWords)
	}

	wc := 0
	for scanner.Scan() {
		wc++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "reading input: %v\n", err)
		return 0
	}

	return wc
}
