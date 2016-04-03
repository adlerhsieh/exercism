package diffsquares

// SquareOfSums takes one argument n,
// sums of all numbers and returns its square.
func SquareOfSums(n int) int {
	var sum int
	for i := 0; i < n+1; i++ {
		sum += i
	}
	return sum * sum
}

// SumOfSquares takes one argument n,
// sqares each and returns the sum.
func SumOfSquares(n int) int {
	var sum int
	for i := 0; i < n+1; i++ {
		sum += i * i
	}
	return sum
}

// Difference returns the difference between
// SquareOfSums and SumOfSquares with argument n.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}
