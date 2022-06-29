package main

import (
	"fmt"
)

func main() {
	var a int
	fmt.Scan(&a)
	fmt.Println(fib(a))
}

func fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return fib(n-1) + fib(n-2)

}
