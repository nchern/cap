// This utility scans lines of a given text. If a line is a heading(e.g. `* Header 1`),
// it checks whether it matechs given regex. If it matches, the heading along with its contents(i.e. chapter)
// will be printed to stdout
package main

import (
	"flag"
	"log"
	"os"

	"github.com/nchern/cap/chapter"
)

// TODO:
// - header should be determined by a configurable pattern
// - ? case insensitive search?

var includeSubChapters = flag.Bool("i", false,
	"If set, all sub-chapters of the matched chapters are also printed out. Subchapter is a chapter with headings of higher levels that the initial one")

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

	p := chapter.NewParser(os.Stdin).IncludeSubChapters(*includeSubChapters)

	must(p.Parse(pattern, os.Stdout))
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
