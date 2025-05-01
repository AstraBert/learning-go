package cryptosquare

import (
	"strings"
	"math"
)

var punctuation = []string{
	"!", "\"", "#", "$", "%", "&", "(", ")", "*", "+", ",", "-", ".", "/", 
	":", ";", "<", "=", ">", "?", "@", "[", "\\", "]", "^", "_", "`", 
	"{", "|", "}", "~", "'",
}

func NoPunctuation(s string) string {
	for _,r := range punctuation {
		s = strings.ReplaceAll(s, r, "")
	}
	return s
}

func FindRectangle(n int) (int, int) {
	if int(math.Sqrt(float64(n)))*int(math.Sqrt(float64(n))) == n {
		return int(math.Sqrt(float64(n))),int(math.Sqrt(float64(n)))
	}
	r := int(math.Sqrt(float64(n)))
	c := r+1
	for i:=1;i>0;i++{
		if r*c >= n && c >= r && (c-r)<=1 {
			return r,c
		} 
		r+=1
		c+=1
	}
	return r,c
}

func Encode(pt string) string {
	pt = NoPunctuation(pt)
	pt = strings.Join("",strings.Fields(pt))
	r,c := FindRectangle(len(pt))	
	// unfinished
}
