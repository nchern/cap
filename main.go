package main

import (
	"log"
	"os"

	"github.com/nchern/cap/chapter"
)

// TODO:
// - add tests
// - header should be determined by a configurable pattern
// - matching with headers should be via regex, not just index
// - option to output matched header with all subheaders

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	pattern := os.Args[1]

	must(chapter.Parse(os.Stdin, pattern, os.Stdout))
}
