[![Go Report Card](https://goreportcard.com/badge/github.com/nchern/cap)](https://goreportcard.com/report/github.com/nchern/cap)
[![Coverage](https://gocover.io/_badge/github.com/nchern/cap)](https://gocover.io/github.com/nchern/cap)

# cap
A command line tool to parse markdown-like formatted texts into chapters and print selected chapters out

> cap. - an abbreviation of capitulum ("chapter").

## Install
```bash
make install # You need go compiler set up.
```

## Usage
```bash
$ cap -h
Usage: cap [FLAGS] [pattern]
  -s	If set, all sub-chapters of the matched chapters are also printed out. Subchapter is a chapter with headings of higher levels that the initial one

This utility scans lines of a given text. If a line is a heading(e.g. '* Header 1'),
it checks whether it matches against given regex <pattern>.
If it matches, the heading along with its contents(i.e. chapter) is printed out.
```
