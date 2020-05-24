package chapter

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

var (
	headerRx = regexp.MustCompile(`^\*+?\s`)
)

// Parser parses text to extract matched headings with their contents
type Parser struct {
	printSubHeaders bool
	r               io.Reader
}

// NewParser returns a new instance of Parser initialized with a given reader
func NewParser(r io.Reader) *Parser {
	return &Parser{
		r: r,
	}
}

// IncludeSubChapters instructs this parser to output all subheadings of matched headings
func (p *Parser) IncludeSubChapters(b bool) *Parser {
	p.printSubHeaders = b
	return p
}

// Parse reads the text from given reader line by line, searches for headings that match given pattern
// and outputs these headings along with their contents to the writer
func (p *Parser) Parse(pattern string, w io.Writer) error {
	type foo struct {
	}

	patternRx, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	currentDepth := 0
	shouldPrint := false

	scanner := bufio.NewScanner(p.r)
	for scanner.Scan() {
		line := scanner.Text()

		header := headerRx.MatchString(line)

		if header {
			depth := getDepth(headerRx.FindStringSubmatch(line)[0])

			matches := patternRx.MatchString(line)
			if !shouldPrint && matches {
				shouldPrint = true
				currentDepth = depth
			} else if shouldPrint && !matches {
				if !(p.printSubHeaders && depth > currentDepth) {
					shouldPrint = false
				}
			}
		}

		if shouldPrint {
			if _, err := fmt.Fprintln(w, scanner.Text()); err != nil {
				return err
			}
		}
	}
	return scanner.Err()
}

func getDepth(header string) int {
	return len(strings.TrimSpace(header))
}
