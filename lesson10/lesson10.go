package main

import (
	"GO01/lesson10/fibonachi"
	"GO01/lesson10/sum"
	"fmt"
)

func main() {
	// file fibonachi
	// fibonachi.Fibonachi()
	res := fibonachi.Fib(15)
	fmt.Println(res)

	res2 := fibonachi.Fib2(15)
	fmt.Println(res2)

	// file sum
	result := sum.Sum(4, 8)
	fmt.Println(result)

}
