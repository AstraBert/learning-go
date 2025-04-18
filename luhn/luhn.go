package luhn

import (
	"strings"
	"strconv"
)

func StringToList(a string) []string {
	b := strings.Split(a, "")
	return b
}

func Reverse(numbers []string) []string {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

func Valid(id string) bool {
	newId := strings.ReplaceAll(id, " ", "")
	if len(newId) <= 1 {
		return false
	}

	splitId := StringToList(newId)
	reversedId := Reverse(splitId)
	checkSum := 0

	for i, x := range reversedId {
		num, err := strconv.Atoi(string(x))
		if err != nil { 
			return false
		}

		if i%2 == 1 {
			num *= 2
			if num > 9 {
				num -= 9
			}
		}
		checkSum += num
	}
	return checkSum%10 == 0
}