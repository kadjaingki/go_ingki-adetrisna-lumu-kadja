package main

import "fmt"

func main() {
	fmt.Println(fibonanci(0))
	fmt.Println(fibonanci(2))
	fmt.Println(fibonanci(9))
	fmt.Println(fibonanci(10))
	fmt.Println(fibonanci(12))
}

func fibonanci(n int) int {
	if n <= 1 {
		return n

	}
	return fibonanci(n-1) + fibonanci(n-2)
}
