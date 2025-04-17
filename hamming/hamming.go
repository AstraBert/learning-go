package hamming

import (
	"errors"
	"strings"
)

var stringLenghtMismatchError = errors.New("Strings length does not match")

func StringToList(a string) []string {
	b := strings.Split(a, "")
	return b
}

func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return 0, stringLenghtMismatchError
	} else {
		hammingDist := 0
		aList := StringToList(a)
		bList := StringToList(b)
		for i := 0; i < len(aList); i++ {
			if aList[i] != bList[i] {
				hammingDist++
			}
		}
		return hammingDist, nil
	}
}
