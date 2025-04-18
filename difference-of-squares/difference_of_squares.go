package diffsquares

func SquareOfSum(n int) int {
	sumNs := 0
	for i:=0; i<=n; i++ {
		sumNs += i
	}
	return sumNs*sumNs
}

func SumOfSquares(n int) int {
	sumNs := 0
	for i:=1; i<=n; i++ {
		sumNs += i*i
	}
	return sumNs
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
