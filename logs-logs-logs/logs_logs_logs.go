package logs

import (
	"strings"
	"unicode/utf8"
)

// Application identifies the application emitting the given log.
func Application(log string) string {
	var app string
	for _, character := range log {
		if character == '❗' {
			return "recommendation"
		} else if character == '🔍' {
			return "search"
		} else if character == '☀' {
			return "weather"
		} else {
			app = "default"
		}
	}
	return app
}

// Replace replaces all occurrences of old with new, returning the modified log
// to the caller.
func Replace(log string, oldRune, newRune rune) string {
	replacedString := strings.Map(func(r rune) rune {
        if r == oldRune {
            return newRune
        }
        return r
    }, log)
	return replacedString
}

// WithinLimit determines whether or not the number of characters in log is
// within the limit.
func WithinLimit(log string, limit int) bool {
	return utf8.RuneCountInString(log) <= limit
}
