package fibonachi

import (
	"fmt"
)

func Fibonachi() {

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

		val = Fib(int(key))
		fmt.Println("Значение Фибоначчи:", val)

		recursionMap[key] = val
		fmt.Println(recursionMap)

	}
}

func Fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)

}

func Fib2(n int) int {
	var f1, f2 int
	f2 = 1
	var i int
	for ; i < n-1; i++ {
		sum := f1 + f2
		f1 = f2
		f2 = sum
	}
	return f2
}
