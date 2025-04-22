package grains

import (
	"errors"
	"math"
)
func Square(number int) (uint64, error) {
	if number < 1 || number > 64 {
		return 0, errors.New("Square number out of interval [1,64]")
	} else {
		return uint64(math.Pow(2, float64(number-1))), nil
	}
}

func Total() uint64 {
	var totalSum uint64
	for i := 0; i <= 63; i++ {
		totalSum += uint64(math.Pow(2, float64(i)))
	}
	return totalSum
}
