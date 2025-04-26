package prime

import (
	"slices"
	"math"
)

func IsPrime(k int64) bool {
	if k <= 1 {
		return false
	} else if k == 2 || k == 3 {
		return true
	} else if k%2 == 0 || k%3 == 0 {
		return false
	} else {
		for i := int64(5); i*i <= k; i = i + 6 {
			if k%i == 0 || k%(i+2) == 0 {
				return false
			}
		}
		return true
	}
}

func Factors(n int64) []int64 {
	if IsPrime(n) {
		return []int64{n}
	}
	factors := []int64{}
	for i:=1; i>0; i++ {
		if n % 2 == 0 {
			n = n / 2
			factors = append(factors, 2)
		} else {
			break
		}
	}
	for i:=int64(3); i < int64(math.Sqrt(float64(n)))+1; i+=2 {
		for j := 1; j > 0; j++ {
			if n % i == 0 {
				n = n / i
				factors = append(factors, i)
			} else {
				break
			}
		}
	}
	if n > 2 {
		factors = append(factors, n)
	} else if n == 2 {
		factors = append(factors, 2)
	}
	slices.Sort(factors)
	return factors
}
