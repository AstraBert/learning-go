package darts

import "math"

func Score(x, y float64) int {
	distanceFromCenter := math.Sqrt(x*x + y*y)
	switch {
	case distanceFromCenter > 10:
		return 0
	case distanceFromCenter <= 10 && distanceFromCenter > 5:
		return 1
	case distanceFromCenter <= 5 && distanceFromCenter > 1:
		return 5
	default:
		return 10
	}
}
