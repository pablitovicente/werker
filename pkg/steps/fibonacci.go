package steps

type Fib struct{}

// The infamously slow recursive Fibonacci algo
func (f Fib) Exec(n uint) uint {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return f.Exec(n-1) + f.Exec(n-2)
	}
}
