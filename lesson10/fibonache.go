package main

import (
	"fmt"
)

func main() {

	recursionMap := map[int]int{
		0: 0,
		1: 1,
	}

	var key int
	var val int

	for {

		fmt.Print("Ввведите число: ")
		_, err := fmt.Scan(&key)

		if err != nil {
			fmt.Println("Вводить можно только числа")
			continue
		}

		if val, ok := recursionMap[key]; ok {
			fmt.Println("Значение Фибоначчи:", val)
			fmt.Println(recursionMap)
			continue

		}

		val = fib(int(key))
		fmt.Println("Значение Фибоначчи:", val)

		recursionMap[key] = val
		fmt.Println(recursionMap)

	}
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
