package calc

func Factorial(n uint) uint {
	if n <= 1 {
		return 1
	}

	return n * Factorial(n-1)
}
