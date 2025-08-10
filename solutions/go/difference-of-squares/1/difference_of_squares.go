package diffsquares

func SquareOfSum(n int) int {
	dt := int(n / 10)
	dr := int(n % 10)

	sum := 0

	for i := (dt*10 + 1); i <= (dt*10 + dr); i++ {
		sum += i
	}

	if dt == 0 {
		return sum * sum
	}

	for i := 0; i < dt; i++ {
		sum += 55 + (100 * i)
	}

	return sum * sum
}

func SumOfSquares(n int) int {
	// TODO: optimize
	sum := 0
	for i := 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

func Difference(n int) int {
	return SquareOfSum(n) - SumOfSquares(n)
}
