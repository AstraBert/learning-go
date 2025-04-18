package isogram

import "strings"

func IsIsogram(word string) bool {
	letterCount := make(map[string]int)
	for i := range word { 
		if letterCount[strings.ToLower(string(word[i]))] >= 1 && strings.ToLower(string(word[i])) != "-" && strings.ToLower(string(word[i])) != " " {
			letterCount[strings.ToLower(string(word[i]))]++
		} else {
			letterCount[strings.ToLower(string(word[i]))] = 1
		}
	}
	for _,v := range letterCount {
		if v > 1 {
			return false
		}
	}
	return true
}
