package mathsutil

import "fmt"

func FibonacciIterative(n int) int {
	if n < 0 {
		return n
	}
	a := 0
	b := 1
	var sum int
	for i := 2; i <= n; i++ {
		sum = a + b
		a = b
		b = sum
	}
	return sum
}

func FibonacciRecursive(n int) int {
	if n < 2 {
		return n
	}
	return FibonacciRecursive(n-1) + FibonacciRecursive(n-1)
}

func FibonacciTests() {
	for i := 0; i <= 10; i++ {
		fmt.Println("fibonacciIterative: ", i, FibonacciIterative(i))
	}
	for i := 0; i <= 10; i++ {
		fmt.Println("fibonacciRecursive: ", i, FibonacciRecursive(i))
	}
}
