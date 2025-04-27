package listops

// IntList is an abstraction of a list of integers which we can define methods on
type IntList []int

func (s IntList) Foldl(fn func(int, int) int, initial int) int {
	for _,x := range s {
		initial = fn(initial,x)
	}
	return initial
}

func (s IntList) Foldr(fn func(int, int) int, initial int) int {
	s = s.Reverse()
	for _,x := range s {
		initial = fn(x,initial)
	}
	return initial
}

func (s IntList) Filter(fn func(int) bool) IntList {
	h := []int{}
	var k IntList
	k = h
	for _,x := range s {
		if fn(x) {
			i := []int{x}
			var j IntList
			j = i
			k = k.Append(j)
		}
	}
	return k	
}

func (s IntList) Length() int {
	return len(s)
}

func (s IntList) Map(fn func(int) int) IntList {
	h := []int{}
	var k IntList
	k = h
	for _,x := range s {
		i := []int{fn(x)}
		var j IntList
		j = i
		k = k.Append(j)
	}
	return k
}

func (s IntList) Reverse() IntList {
	h := []int{}
	var k IntList
	k = h
	sLen := s.Length()
	for i:=sLen - 1;i >= 0; i-- {
		m := []int{s[i]}
		var j IntList
		j = m
		k = k.Append(j)
	} 
	return k
}

func (s IntList) Append(lst IntList) IntList {
	var l IntList
	k := []int{}
	for _,x := range s {
		k = append(k,x)
	}
	for _,x := range lst {
		k = append(k,x)
	}
	l = k
	return l
}

func (s IntList) Concat(lists []IntList) IntList {
	for _,x := range lists {
		s = s.Append(x)
	}
	return s
}
