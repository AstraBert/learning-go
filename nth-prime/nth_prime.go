package prime

import "errors" // Keep the import

func IsPrime(k int) bool {
	if k <= 1 {
		return false
	} else if k == 2 || k == 3 {
		return true
	} else if k%2 == 0 || k%3 == 0 {
		return false
	} else {
		for i := 5; i*i <= k; i = i + 6 {
			if k%i == 0 || k%(i+2) == 0 {
				return false
			}
		}
		return true
	}
}

// Nth returns the nth prime number. An error must be returned if the nth prime number can't be calculated ('n' is equal or less than zero)
func Nth(n int) (int, error) {
	if n < 1 {
		return 0, errors.New("n must be a positive integer")
	}

	count := 0
	num := 2

	for count < n {
		if IsPrime(num) {
			count++ 
		}
		num++
	}
	return num - 1, nil
}