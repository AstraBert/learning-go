package sorting

import (
	"fmt"
	"strconv"
)


// DescribeNumber should return a string describing the number.
func DescribeNumber(f float64) string {
	return fmt.Sprintf("This is the number %.1f", f)
}

type Number struct {
	n int
}

func (i Number) Number() int {
	return i.n
}

type NumberBox interface {
	Number() int
}

// DescribeNumberBox should return a string describing the NumberBox.
func DescribeNumberBox(nb NumberBox) string {
	var floatNumber float64 = float64(nb.Number())
	return fmt.Sprintf("This is a box containing the number %.1f", floatNumber)
}

type FancyNumber struct {
	n string
}

func (i FancyNumber) Value() string {
	return i.n
}

type FancyNumberBox interface {
	Value() string
}

// ExtractFancyNumber should return the integer value for a FancyNumber
// and 0 if any other FancyNumberBox is supplied.
func ExtractFancyNumber(fnb FancyNumberBox) int {
	switch fnb.(type) {
	case FancyNumber:
		nb := fnb.Value()
		i, _ := strconv.Atoi(nb)
		return i
	default:
		return 0		
	}
}

// DescribeFancyNumberBox should return a string describing the FancyNumberBox.
func DescribeFancyNumberBox(fnb FancyNumberBox) string {
	nbf := float64(ExtractFancyNumber(fnb))
	return fmt.Sprintf("This is a fancy box containing the number %.1f", nbf)
}

// DescribeAnything should return a string describing whatever it contains.
func DescribeAnything(i interface{}) string {
	nf, ok0 := i.(int)
	if ok0 {	
		return DescribeNumber(float64(nf))
	} else {
		nb, ok := i.(float64)
		if ok {
			return DescribeNumber(nb)
		} else {
			str, ok1 := i.(FancyNumberBox)
			if ok1 {
				return DescribeFancyNumberBox(str)
			} else {
				ni, ok2 := i.(NumberBox)
				if ok2 {
					return DescribeNumberBox(ni)
				} else {
					return "Return to sender"
				}
			}
		}
	}
}