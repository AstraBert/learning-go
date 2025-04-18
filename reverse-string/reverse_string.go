package reverse

import "strings"

func StringToList(a string) []string {
	b := strings.Split(a, "")
	return b
}

func Reverse(input string) string {
	splitInput := StringToList(input)
	reversedInput := []string{}
	for i := len(splitInput) - 1; i >= 0; i-- {
		reversedInput = append(reversedInput, splitInput[i])
	}
	return strings.Join(reversedInput, "")
}
