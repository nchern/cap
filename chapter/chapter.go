package chapter

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

// Parse reads the text from given reader line by line, searches for headers that match given pattern
// and outputs these headers along with their contents to the writer
func Parse(r io.Reader, pattern string, w io.Writer) error {
	shouldPrint := false

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())

		header := isHeader(line)

		if header {
			matched := strings.Index(line, pattern) > -1
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

func isHeader(line string) bool {
	headerPattern := "*"
	return strings.HasPrefix(line, headerPattern) && !strings.HasSuffix(line, headerPattern)
}
