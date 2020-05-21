package chapter

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
)

var (
	headerRx = regexp.MustCompile(`^\*+?\s`)
)

// Parse reads the text from given reader line by line, searches for headers that match given pattern
// and outputs these headers along with their contents to the writer
func Parse(r io.Reader, pattern string, w io.Writer) error {
	patternRx, err := regexp.Compile(pattern)
	if err != nil {
		return err
	}

	shouldPrint := false

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := scanner.Text()

		header := headerRx.MatchString(line)

		if header {
			matched := patternRx.MatchString(line)
			if !shouldPrint && matched {
				shouldPrint = true
			} else if shouldPrint && !matched {
				shouldPrint = false
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
