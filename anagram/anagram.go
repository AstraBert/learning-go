package anagram

import (
	"strings"
	"slices"
)

func SplitString(s string) []string {
	return strings.Split(s, "")
}

func Detect(subject string, candidates []string) []string {
	anagrams := []string{}
	splitSub := SplitString(strings.ToLower(subject))
	for _,x := range candidates {
		if len(x) == len(subject) && strings.ToLower(x) != strings.ToLower(subject) {
			splitX := SplitString(strings.ToLower(x))
			truths := []bool{}
			for _,c := range splitX {
				if slices.Contains(splitSub, c) && strings.Count(strings.ToLower(subject), c) == strings.Count(strings.ToLower(x), c) {
					truths = append(truths, true)
				}
			}
			if len(truths) == len(splitX) {
				anagrams = append(anagrams, x)
			}
		} 
	}
	return anagrams
}
