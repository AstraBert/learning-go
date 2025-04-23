package wordcount

import (
	"strings"
)

var punctuation = []string{
	"!", "\"", "#", "$", "%", "&", "(", ")", "*", "+", ",", "-", ".", "/", 
	":", ";", "<", "=", ">", "?", "@", "[", "\\", "]", "^", "_", "`", 
	"{", "|", "}", "~",
}

type Frequency map[string]int

func WordCount(phrase string) Frequency {
	str := strings.ToLower(phrase)
	str = strings.ReplaceAll(str," '"," ")
	str = strings.ReplaceAll(str,"' "," ")
	str, _ = strings.CutSuffix(str, "'")
	str, _ = strings.CutPrefix(str, "'")
	for _,x := range punctuation {
		str = strings.ReplaceAll(str, x, " ")
	}
	words := strings.Fields(str)
	var k Frequency
	k = map[string]int{}
	for _,x := range words {
		count := 0
		for _,y := range words {
			if x == y {
				count++
			}
		}
		k[x] = count
	}
	return k
}
