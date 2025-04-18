package dna

import (
	"unicode"
	"errors"
)

var InvalidDnaStrandError = errors.New("Invalid DNA strand: a DNA strand must only contain 'A', 'T', 'C' or 'G'")

// Histogram is a mapping from nucleotide to its count in given DNA.
// Choose a suitable data type.
type Histogram map[rune]int

// DNA is a list of nucleotides. Choose a suitable data type.
type DNA string

func (d DNA) Counts() (Histogram, error) {
	nucleotides := Histogram{'A': 0, 'T': 0, 'C': 0, 'G': 0}
	for _,x := range d {
		if unicode.To(unicode.UpperCase, x) != 'A' && unicode.To(unicode.UpperCase, x) != 'C' && unicode.To(unicode.UpperCase, x) != 'T' && unicode.To(unicode.UpperCase, x) != 'G' {
			return nil, InvalidDnaStrandError
		} else {
			nucleotides[x] += 1
		}
	}
	return nucleotides, nil
}

