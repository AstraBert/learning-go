package sieve

import "slices"

func Sieve(limit int) []int {
	if limit < 2 {
		return []int{}
	} else if limit == 2 {
		return []int{2}
	} else {
		sieve := map[int]bool{}
		nums := []int{}
		for i:=2; i<=limit; i++ {
			sieve[i] = true
			nums = append(nums, i)
		}
		for _,x := range nums {
			if sieve[x] {
				for _,k := range nums {
					if x != k && k % x == 0 {
						sieve[k] = false
					}
				}
			} else {
				continue
			}
		}
		primes := []int{}
		for k,v := range sieve {
			if v {
				primes = append(primes, k)
			}
		}
		slices.Sort(primes)
		return primes
	}
}
