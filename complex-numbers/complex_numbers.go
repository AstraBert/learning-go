package complexnumbers

import "math"

// Define the Number type here.
type Number struct {
	a float64
	b float64
}

func (n Number) Real() float64 {
	return n.a
}

func (n Number) Imaginary() float64 {
	return n.b
}

func (n1 Number) Add(n2 Number) Number {
	n3 := Number{a: n1.Real()+n2.Real(), b: n1.Imaginary()+n2.Imaginary()}
	return n3
}

func (n1 Number) Subtract(n2 Number) Number {
	n3 := Number{a: n1.Real()-n2.Real(), b: n1.Imaginary()-n2.Imaginary()}
	return n3
}

func (n1 Number) Multiply(n2 Number) Number {
	n3 := Number{a: n1.Real()*n2.Real()-n1.Imaginary()*n2.Imaginary(), b: n1.Imaginary()*n2.Real()+n1.Real()*n2.Imaginary()}
	return n3
}

func (n Number) Times(factor float64) Number {
	n3 := Number{a: n.Real()*factor, b: n.Imaginary()*factor}
	return n3
}

func (n1 Number) Divide(n2 Number) Number {
	a := n1.Real()*n2.Real()+n1.Imaginary()*n2.Imaginary()
	b := n1.Imaginary()*n2.Real()-n1.Real()*n2.Imaginary()
	n4 := Number{a: a/(math.Pow(n2.Real(),2)+math.Pow(n2.Imaginary(),2)), b: b/(math.Pow(n2.Real(),2)+math.Pow(n2.Imaginary(),2))}
	return n4
}

func (n Number) Conjugate() Number {
	n3 := Number{a: n.Real(), b: -n.Imaginary()}
	return n3
}

func (n Number) Abs() float64 {
	return math.Sqrt(math.Pow(n.a,2)+math.Pow(n.b,2))
}

func (n Number) Exp() Number {
	a := math.Exp(n.Real())*math.Cos(n.Imaginary())
	b := math.Exp(n.Real())*math.Sin(n.Imaginary())
	n4 := Number{a: a, b: b}
	return n4	
}
