package resistorcolor

// Colors returns the list of all colors.
func Colors() []string {
	colorsList := []string{"black", "brown","red","orange","yellow","green","blue","violet","grey","white"}
	return colorsList
}

// ColorCode returns the resistance value of the given color.
func ColorCode(color string) int {
	colorsMap := make(map[string]int)
	colors := Colors()
	for i,x := range colors {
		colorsMap[x] = i
	}
	return colorsMap[color]
}
