package isbn

import (
	"strings"
	"strconv"
)

func RemoveHyphens(c string) string {
	c1 := strings.ReplaceAll(c, "-", "")
	return c1
}

func SplitToNumbers(c string) []int {
	splitCode := strings.Split(c, "")
	numbers := []int{}
	for _,x := range splitCode {
		if x != "X" {
			num, _ := strconv.Atoi(x)
			numbers = append(numbers, num)
		} else {
			numbers = append(numbers, 10)
		}
	}
	return numbers
}

func IsCompliant(numbers []int) bool {
	checkSum := 0
	for i,x := range numbers {
		checkSum += x*(10-i)
	}
	remainder := checkSum % 11
	if remainder == 0 {
		return true
	} else {
		return false
	}
}

func IsValidISBN(isbn string) bool {
	if strings.ContainsAny(strings.ToLower(isbn), "abcdefghijklmnopqrstuvwyz") {
		return false
	}
	noHyphens := RemoveHyphens(isbn)
	if len(noHyphens) != 10 {
		return false
	}
	if !strings.HasSuffix(noHyphens, "X") && strings.ContainsAny(noHyphens, "X") {
		return false
	}
	nums := SplitToNumbers(noHyphens)
	return IsCompliant(nums)
}
