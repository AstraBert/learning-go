package etl

import "strings"

func Transform(in map[int][]string) map[string]int {
	mappedIn := make(map[string]int)
	for k,v := range in {
		for _,x := range v {
			mappedIn[strings.ToLower(x)] = k
		}
	}
	return mappedIn
}
