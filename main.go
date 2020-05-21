package main

import (
	"flag"
	"log"
	"os"

	"github.com/nchern/cap/chapter"
)

// TODO:
// - header should be determined by a configurable pattern
// - matching with headers should be via regex, not just index
// - option to output matched header with all subheaders
// - ? case insensitive search?

func init() {
	log.SetFlags(0)
	flag.Parse()

	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(2)
	}
}

func main() {
	pattern := flag.Arg(0)

	must(chapter.Parse(os.Stdin, pattern, os.Stdout))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
