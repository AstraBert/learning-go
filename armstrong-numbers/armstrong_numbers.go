package armstrong

import (
	"strings"
	"strconv"
	"math"
)

func toList(n int) []int {
	s := strconv.Itoa(n)
	splitS := strings.Split(s, "")
	splitN := []int{}
	for _,x := range splitS {
		num, _ := strconv.Atoi(string(x))
		splitN = append(splitN, num)
	}
	return splitN
}

func IsNumber(n int) bool {
	numList := toList(n)
	sumN := 0
	for _,n := range numList {
		sumN += int(math.Pow(float64(n), float64(len(numList))))
	}
	return sumN == n
}
