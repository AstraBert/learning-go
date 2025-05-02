package cryptosquare

import (
	"math"
	"strings"
	"unicode"
)

func NoPunctuation(s string) string {
	var b strings.Builder
	for _, r := range s {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			b.WriteRune(unicode.ToLower(r))
		}
	}
	return b.String()
}

func FindRectangle(n int) (int, int) {
	sqrt := int(math.Sqrt(float64(n)))
	for r := sqrt; r <= n; r++ {
		c := (n + r - 1) / r // ceiling
		if c >= r && c-r <= 1 {
			return r, c
		}
	}
	return 0, 0 // Should never happen
}

func Encode(pt string) string {
	pt = NoPunctuation(pt)
	n := len(pt)
	if n == 0 {
		return ""
	}

	r, c := FindRectangle(n)
	// Fill rows
	rows := make([]string, c)
	for i := 0; i < n; i++ {
		rows[i/c] += string(pt[i])
	}
	// Make sure all rows are exactly length c
	for i := range rows {
		if len(rows[i]) < c {
			rows[i] += strings.Repeat(" ", c-len(rows[i]))
		}
	}

	// Transpose to columns
	cols := make([]string, c)
	for col := 0; col < c; col++ {
		for row := 0; row < r; row++ {
			cols[col] += string(rows[row][col])
		}
	}

	// Join with space
	return strings.Join(cols, " ")
}
