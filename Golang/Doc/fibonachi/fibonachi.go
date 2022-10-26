package main

import "fmt"

func main() {
	fmt.Println("work")
	fmt.Println(Fib(4))
	fmt.Println(Fib2(5))
}

// Fib функция фибоначи с рекурсией
func Fib(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	return Fib(n-1) + Fib(n-2)

}

// Fib2 функция фибоначи
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
