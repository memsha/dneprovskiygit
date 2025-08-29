package diffsquares

func SquareOfSum(n int) int {
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i
	}
	return sum * sum

}

func SumOfSquares(n int) int {
	square := 0
	for i := 1; i <= n; i++ {
		square += i * i
	}
	return square
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
