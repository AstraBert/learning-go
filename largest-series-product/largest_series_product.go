package lsproduct

import (
	"strings"
	"strconv"
	"errors"
	"slices"
)

func DivideBySpan(s string, d int) []string {
	r := []string{}
	if len(s) == d {
		return []string{s}
	}
	for i,_ := range s {
		if i+d-1 < len(s) {
			r = append(r, s[i:i+d])
		}
	}
	return r
}

func LargestSeriesProduct(digits string, span int) (int64, error) {
	if  span > len(digits) {
		return 0, errors.New("span must be smaller than string length")
	} 
	if span < 0 {
		return 0, errors.New("span must not be negative")
	}
	if strings.ContainsAny(digits, "abcdefghijklmnopqrstuvwxyz") {
		return 0, errors.New("digits input must only contain digits")
	}
	ds := DivideBySpan(digits, span)
	prods := []int64{}
	for _,x := range ds {
		ls := strings.Split(x, "")
		prod := 1
		for _,l := range ls {
			num,_ := strconv.Atoi(string(l))
			prod = prod*num
		}
		prods = append(prods, int64(prod))
	}
	return slices.Max(prods), nil
}
