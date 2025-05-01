package bottlesong

import (
	"fmt"
	"strings"
)

var num2Literal = map[int]string{
	0: "no",
	1:  "one",
	2:  "two",
	3:  "three",
	4:  "four",
	5:  "five",
	6:  "six",
	7:  "seven",
	8:  "eight",
	9:  "nine",
	10: "ten",
}

func Recite(startBottles, takeDown int) []string {
	listofStrings := []string{}
	for i:=startBottles; i > (startBottles-takeDown); i-- {
		if num2Literal[i] != "one" {
			s0 := string(strings.ToUpper(string([]byte{num2Literal[i][0]}))+num2Literal[i][1:])
			s1 := fmt.Sprintf("%s green bottles hanging on the wall,", s0)
			s2 := fmt.Sprintf("%s green bottles hanging on the wall,", s0)
			s3 := fmt.Sprintf("And if one green bottle should accidentally fall,")
			if num2Literal[i-1] != "one" {
				s4 := fmt.Sprintf("There'll be %s green bottles hanging on the wall.", num2Literal[i-1])
				if i-1 > (startBottles-takeDown) {
					listofStrings = append(listofStrings, s1, s2, s3, s4, "")
				} else {
					listofStrings = append(listofStrings, s1, s2, s3, s4)
				}
			} else {
				s4 := fmt.Sprintf("There'll be %s green bottle hanging on the wall.", num2Literal[i-1])
				if i-1 > (startBottles-takeDown) {
					listofStrings = append(listofStrings, s1, s2, s3, s4, "")
				} else {
					listofStrings = append(listofStrings, s1, s2, s3, s4)
				}
			}
		} else {
			s0 := string(strings.ToUpper(string([]byte{num2Literal[i][0]}))+num2Literal[i][1:])
			s1 := fmt.Sprintf("%s green bottle hanging on the wall,", s0)
			s2 := fmt.Sprintf("%s green bottle hanging on the wall,", s0)
			s3 := fmt.Sprintf("And if one green bottle should accidentally fall,")
			s4 := fmt.Sprintf("There'll be %s green bottles hanging on the wall.", num2Literal[i-1])
			listofStrings = append(listofStrings, s1, s2, s3, s4)
		}
	}
	return listofStrings
}
