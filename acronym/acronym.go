// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package acronym should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package acronym

import "strings"

var punctuation = []string{
	"!", "\"", "#", "$", "%", "&", "(", ")", "*", "+", ",", ".", "/", 
	":", ";", "<", "=", ">", "?", "@", "[", "\\", "]", "^", "_", "`", 
	"{", "|", "}", "~", "'",
}

func NoPunctuation(s string) string {
	for _,r := range punctuation {
		s = strings.ReplaceAll(s, r, "")
	}
	return s
}

// Abbreviate should have a comment documenting it.
func Abbreviate(s string) string {
	s = NoPunctuation(s)
	s = strings.ToTitle(s)
	s = strings.ReplaceAll(s, "-", " ")
	fs := strings.Fields(s)
	acronym := ""
	for _,x := range fs {
		acronym += string([]byte{x[0]})
	}
	return acronym
}
