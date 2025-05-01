package wordy

import (
	"slices"
	"strconv"
	"strings"
)

var allowedOperations = []string{
	"plus",
	"minus",
	"multiplied",
	"divided",
}

func Answer(question string) (int, bool) {
	question = strings.ReplaceAll(question,"?","")
	comps := strings.Split(question," ")
	nums := []int{}
	ops := []string{}
	for _,x := range comps {
		num, err := strconv.Atoi(x)
		if err == nil {
			nums = append(nums, num)
		} else {
			if slices.Contains(allowedOperations, x) {
				ops = append(ops, x)
			}
		}
		if len(nums) > len(ops)+1 {
			return 0, false
		} 
		if len(ops) > len(nums) {
			return 0, false
		}
	}
	if len(nums) == 1 && len(ops) == 0 && len(comps) == 3 {
		return nums[0], true
	}
	if len(nums) == 0 || len(ops) == 0 || len(ops) != (len(nums)-1) {
		return 0, false
	}
	result := nums[0]
	for i,x := range ops {
		switch x {
		case "plus":
			result = result+nums[i+1]
		case "minus":
			result = result-nums[i+1]
		case "multiplied":
			result = result*nums[i+1]
		case "divided":
			result = result/nums[i+1]
		} 
	}
	return result, true
}
