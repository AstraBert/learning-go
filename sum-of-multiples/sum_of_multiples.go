package summultiples

import "slices"

func MultiplesUpTo(n, limit int) []int {
    if n <= 0 || limit <= 0 {
        return []int{}
    }
    numMultiples := limit / n
    multiples := make([]int, 0, numMultiples)
    for s := n; s < limit; s += n {
        multiples = append(multiples, s)
    }
    return multiples
}

func SumList(l []int) int {
	s:=0
	for _,x := range l {
		s+=x
	}
	return s
}

func SumMultiples(limit int, divisors ...int) int {
	listOfMultiples := []int{}
	for _,d := range divisors {
		listOfMultiples = slices.Concat(listOfMultiples, MultiplesUpTo(d,limit))
	}
	uniqueMultiples := []int{}
	for _,m := range listOfMultiples {
		if !slices.Contains(uniqueMultiples, m) {
			 uniqueMultiples = append(uniqueMultiples, m)
		} 
	} 
	return SumList(uniqueMultiples)
}
