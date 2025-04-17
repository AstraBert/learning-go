package collatzconjecture

import "errors"

var nonPositiveNumberError = errors.New("You provided a non positive number")

func CollatzConjecture(n int) (int, error) {
	if n <= 0 {
		return 0, nonPositiveNumberError
	} else {
		if n == 1 {
			return 0, nil
		} else {
			steps := 0
			for i := 0; i > -1; i++ {
				if n % 2 == 0 {
					n = n/2
					steps++
				} else {
					n = 3*n + 1
					steps++
				}
				if n == 1 {
					break
				}
			}
			return steps, nil
		}
	}
}
