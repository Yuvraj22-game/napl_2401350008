package test1

func Factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * Factorial(n-1)
}
func Power(m, n int) int {
	if n == 0 {
		return 1
	}
	return m * Power(m, n-1)
}
