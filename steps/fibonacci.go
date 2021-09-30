package steps

// The infamously slow recursive Fibonacci algo
func Fib(n uint) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return Fib(n-1) + Fib(n-2)
	}
}
