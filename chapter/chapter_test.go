package chapter

import (
	"bytes"
	"strings"
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
			text([]string{
				"* main",
				"this is main",
				"** subheader 1",
				"*Bold text*",
				"foo bar",
				"** something else",
				"foo boo",
				"** subheader 2",
				"fuzz buzz",
				"*** deeper",
				"hello world",
				"* main 2",
				"* Subheader 3",
				"hello Subheader"}),
			text([]string{
				"** subheader 1",
				"*Bold text*",
				"foo bar",
				"** subheader 2",
				"fuzz buzz",
				""})},
		{"sequential headers",
			"header",
			text([]string{
				"* main",
				"this is main",
				"** header 1",
				"*Bold text* with not bold!",
				"foo bar",
				"** header 2",
				"fuzz buzz",
				"*** deeper",
				"hello world",
				"* main 2"}),
			text([]string{
				"** header 1",
				"*Bold text* with not bold!",
				"foo bar",
				"** header 2",
				"fuzz buzz",
				""})},
		{"pattern is regex",
			`header \d$`,
			text([]string{
				"* main",
				"this is main",
				"** header 1",
				"*Bold text* with not bold!",
				"foo bar",
				"** header 2",
				"fuzz buzz",
				"** header A",
				"hello world",
				"* header B",
				"barrr"}),
			text([]string{
				"** header 1",
				"*Bold text* with not bold!",
				"foo bar",
				"** header 2",
				"fuzz buzz",
				""})},
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

func TestShouldParseFailIfPrefixMakesYieldsRegexp(t *testing.T) {
	var actualBuf bytes.Buffer

	underTest := NewParser(bytes.NewBufferString("* Header")).SetPrefix("()")
	assert.Error(t, underTest.Parse("Header.*$", &actualBuf))
}

func TestShouldParseWithDiffentHeaderMarker(t *testing.T) {
	given := text([]string{
		"intro",
		"# main",
		"foo bar",
		"# second A",
		"fuzz buzz",
		"## level 2",
		"hello",
		"### level 3",
		"test",
		"## level 2 another",
		"hello 2",
		"# second B",
		"lala lala",
		"# third",
		"hey"})

	expected := text([]string{
		"# second A",
		"fuzz buzz",
		"## level 2",
		"hello",
		"### level 3",
		"test",
		"## level 2 another",
		"hello 2",
		"# second B",
		"lala lala",
		""})

	var actualBuf bytes.Buffer

	underTest := NewParser(bytes.NewBufferString(given)).
		IncludeSubChapters(true).
		SetPrefix("#")

	assert.NoError(t, underTest.Parse("second.*$", &actualBuf))
	assert.Equal(t, expected, actualBuf.String())
}

func TestShouldParseAndIgnoreCase(t *testing.T) {
	given := text([]string{
		"* main",
		"hello",
		"** HeADer 1",
		"foobar",
		"HEADER 2",
		"fuzz buzz",
		"** hdr",
		"bar"})

	expected := text([]string{
		"** HeADer 1",
		"foobar",
		"HEADER 2",
		"fuzz buzz",
		""})

	var actualBuf bytes.Buffer

	underTest := NewParser(bytes.NewBufferString(given)).IgnoreCase(true)

	assert.NoError(t, underTest.Parse("head", &actualBuf))
	assert.Equal(t, expected, actualBuf.String())
}

func text(lines []string) string {
	return strings.Join(lines, "\n")
}
