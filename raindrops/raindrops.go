package raindrops

import "strconv"

func Convert(number int) string {
	result := ""
	sounds := make(map[int]string)
	sounds[3] = "Pling"
	sounds[5] = "Plang"
	sounds[7] = "Plong"
	divideBy := []int{3,5,7}
	for _,i := range divideBy {
		if number % i == 0{
			result += sounds[i]
		}
	}
	if result == "" {
		return strconv.Itoa(number)
	} else {
		return result
	}
}
