package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/nchern/cap/chapter"
)

const help = `
This utility scans lines of a given text. If a line is a heading(e.g. '* Header 1'),
it checks whether it matechs against given regex <pattern>. If it matches, the heading along with its contents(i.e. chapter)
will be printed to stdout`

var (
	includeSubChapters = flag.Bool("s", false,
		"If set, all sub-chapters of the matched chapters are also printed out. Subchapter is a chapter with headings of higher levels that the initial one")

	ignoreCase = flag.Bool("i", false, "Perform case insensitive heading matching")
)

func init() {
	log.SetFlags(0)
	flag.Usage = usage

	flag.Parse()

	if len(flag.Args()) < 1 {
		flag.Usage()
		os.Exit(2)
	}
}

// TODO:
// - specify file to parse, not only stdin
// - (?) header should be determined by a configurable pattern

func main() {
	pattern := flag.Arg(0)

	p := chapter.NewParser(os.Stdin).
		IgnoreCase(*ignoreCase).
		IncludeSubChapters(*includeSubChapters)

	must(p.Parse(pattern, os.Stdout))
}

func usage() {
	out := flag.CommandLine.Output()
	fmt.Fprintf(out, "Usage: %s [flags] [pattern]\n", os.Args[0])
	flag.PrintDefaults()
	fmt.Fprintln(out, help)
}

func must(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
