package chapter

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShouldParse(t *testing.T) {
	var tests = []struct {
		name         string
		givenPattern string
		givenText    string
		expected     string
	}{
		{"empty text", "foo", "", ""},
		{"sample_1",
			"subheader",
			`
* main
this is main
** subheader 1
*Bold text*
foo bar
** something else
foo boo
** subheader 2
fuzz buzz
*** deeper
hello world
* main 2`,
			`** subheader 1
*Bold text*
foo bar
** subheader 2
fuzz buzz
`},
		{"sequential headers",
			"header",
			`
* main
this is main
** header 1
*Bold text* with not bold!
foo bar
** header 2
fuzz buzz
*** deeper
hello world
* main 2`,
			`** header 1
*Bold text* with not bold!
foo bar
** header 2
fuzz buzz
`},
		{"patter is regex",
			`header \d$`,
			`
* main
this is main
** header 1
*Bold text* with not bold!
foo bar
** header 2
fuzz buzz
** header A
hello world
* header B
barrr`,
			`** header 1
*Bold text* with not bold!
foo bar
** header 2
fuzz buzz
`},
	}
	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			var actual bytes.Buffer
			assert.NoError(t, NewParser(bytes.NewBufferString(tt.givenText)).Parse(tt.givenPattern, &actual))
			assert.Equal(t, tt.expected, actual.String())
		})
	}
}
