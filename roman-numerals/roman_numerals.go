package romannumerals

import (
	"errors"
	"strconv"
	"strings"
)

func ToIntList(n int) []int {
	splitStr := strings.Split(strconv.Itoa(n), "")
	numList := []int{}
	for _,x := range splitStr {
		num, _ := strconv.Atoi(string(x))
		numList = append(numList, num)
	}
	return numList
}

func Reverse(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}
	return numbers
}

var unitsToRoman =  map[int]string{
	0: "", 
	1: "I",
	2: "II",
	3: "III",
	4: "IV",
	5: "V",
	6: "VI",
	7: "VII",
	8: "VIII",
	9: "IX",
}

var tenthsToRoman =  map[int]string{
	0: "", 
	1: "X",
	2: "XX",
	3: "XXX",
	4: "XL",
	5: "L",
	6: "LX",
	7: "LXX",
	8: "LXXX",
	9: "XC",
}

var hundsToRoman =  map[int]string{
	0: "", 
	1: "C",
	2: "CC",
	3: "CCC",
	4: "CD",
	5: "D",
	6: "DC",
	7: "DCC",
	8: "DCCC",
	9: "CM",
}

func CreateRomanNumeral(numList []int) string {
	newNums := Reverse(numList)
	if len(newNums) == 4 {
		units := unitsToRoman[newNums[0]]
		tenths := tenthsToRoman[newNums[1]]
		hunds := hundsToRoman[newNums[2]]
		thous := strings.Repeat("M", newNums[3])
		return thous+hunds+tenths+units
	} else if len(newNums) == 3 {
		units := unitsToRoman[newNums[0]]
		tenths := tenthsToRoman[newNums[1]]
		hunds := hundsToRoman[newNums[2]]
		return hunds+tenths+units
	} else if len(newNums) == 2 {
		units := unitsToRoman[newNums[0]]
		tenths := tenthsToRoman[newNums[1]]
		return tenths+units
	} else {
		return unitsToRoman[newNums[0]]
	}
}

func ToRomanNumeral(input int) (string, error) {
	if input <= 0 || input > 3999 {
		return "", errors.New("Invalid input number")
	} else {
		numList := ToIntList(input)
		return CreateRomanNumeral(numList), nil
	}
}
