package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	wc := count(os.Stdin)
	fmt.Println(wc)
}

func count(r io.Reader) int {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)

	wc := 0
	for scanner.Scan() {
		wc++
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "reading input: %v\n", err)
	}

	return wc
}
