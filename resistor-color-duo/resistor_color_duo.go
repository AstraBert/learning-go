package resistorcolorduo

import "strconv"

var colorMap = map[string]string{
	"black":  "0",
	"brown":  "1",
	"red":    "2",
	"orange": "3",
	"yellow": "4",
	"green":  "5",
	"blue":   "6",
	"violet": "7",
	"grey":   "8",
	"white":  "9",
}

// Value should return the resistance value of a resistor with a given colors.
func Value(colors []string) int {
	intString := ""
	for i:=0; i<2; i++ {
		intString+=colorMap[colors[i]]
	}
	num,_ := strconv.Atoi(intString)
	return num
}
