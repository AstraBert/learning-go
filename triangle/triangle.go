// Package triangle, starting from three float numbers representing the sides, classifies a three-sided shape into a triangle categories (if it's a triangle)
package triangle


type Kind int

const (
    NaT = iota
    Equ = iota
    Iso = iota
    Sca = iota
)

func IsATriangle(a,b,c float64) bool {
	if a <= 0 || b <= 0 || c <= 0 {
		return false
	}
	return (a+b)>=c && (a+c)>=b && (b+c)>=a
}

func IsIsoscele(a,b,c float64) bool {
	return a == b && a != c || a == c && a != b || b == c && b != a
}

func IsEquilater(a,b,c float64) bool {
	return a == b && b == c
}

// KindFromSides returns the kind of triangle based on the lengths of its sides
func KindFromSides(a, b, c float64) Kind {
	var k Kind
	if !IsATriangle(a,b,c) {
		k = NaT
	} else {
		if IsIsoscele(a,b,c) {
			k = Iso
		} else if IsEquilater(a,b,c) {
			k = Equ
		} else {
			k = Sca
		}
	}
	return k
}
