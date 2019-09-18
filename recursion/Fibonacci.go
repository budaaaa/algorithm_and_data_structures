package recursion

func Fibonacci(n int) int {
	if n > 1 {
		return Fibonacci(n-2) + Fibonacci(n-1)
	}

	return n
}

func FibonacciMemo(n int64) int64 {

}