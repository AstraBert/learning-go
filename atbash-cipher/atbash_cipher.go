package atbash

import "strings"

var letter2Index = map[string]int{
	"a": 1,
	"b": 2,
	"c": 3,
	"d": 4,
	"e": 5,
	"f": 6,
	"g": 7,
	"h": 8,
	"i": 9,
	"j": 10,
	"k": 11,
	"l": 12,
	"m": 13,
	"n": 14,
	"o": 15,
	"p": 16,
	"q": 17,
	"r": 18,
	"s": 19,
	"t": 20,
	"u": 21,
	"v": 22,
	"w": 23,
	"x": 24,
	"y": 25,
	"z": 26,
}

var punctuation = []string{
	"!", "\"", "#", "$", "%", "&", "(", ")", "*", "+", ",", "-", ".", "/", 
	":", ";", "<", "=", ">", "?", "@", "[", "\\", "]", "^", "_", "`", 
	"{", "|", "}", "~", "'",
}

var index2Letter = map[int]string{1: "a", 2: "b", 3: "c", 4: "d", 5: "e", 6: "f", 7: "g", 8: "h", 9: "i", 10: "j", 11: "k", 12: "l", 13: "m", 14: "n", 15: "o", 16: "p", 17: "q", 18: "r", 19: "s", 20: "t", 21: "u", 22: "v", 23: "w", 24: "x", 25: "y", 26: "z"}    

func NoPunctuation(s string) string {
	for _,r := range punctuation {
		s = strings.ReplaceAll(s, r, "")
	}
	return s
}

func DivideString(s string, r int) string {
    if r <= 0 {
        return s
    }
    if len(s) < r {
        return s
    }
	sl := []string{}
	for i := 0; i < len(s); i += r {
		end := i + r
		if end > len(s) {
			end = len(s)
		}
		sl = append(sl, s[i:end])
	}
	return strings.Join(sl, " ")
}

func Atbash(s string) string {
	s = NoPunctuation(s)
	s = strings.ReplaceAll(s," ", "")
	s = strings.ToLower(s)
	splitS := strings.Split(s, "")
	cyphered := ""
	for _,r := range splitS {
		if !strings.ContainsAny(r,"abcdefghijklmnopqrstuvwxyz") {
			cyphered += r
		} else {
			cyphered += index2Letter[27-letter2Index[r]]
		}
	}
	return DivideString(cyphered, 5)
}
