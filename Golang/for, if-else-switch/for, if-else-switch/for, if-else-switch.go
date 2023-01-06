package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	/*
		var command = 3

		if command == 6 {
			fmt.Println(6)
		}
		if command == 7 {
			fmt.Println(7)
		}
		if command == 5 {
			fmt.Println(5)
		} else {
			fmt.Println("ничего не подходит")
		}
	*/

	/*
		var a, b = 5, 3
		var command string
		command = "+"

		switch command {
		case "+":
			fmt.Println(a + b)
		case "-":
			fmt.Println(a - b)
		case "*":
			fmt.Println(a * b)
		case "/":
			fmt.Println(a / b)
		default:
			fmt.Println("Неверная команда")

		}
	*/
	/*
		var count = 10
		for count > 0 {
			fmt.Println(count)
			time.Sleep(time.Second)
			count--
		}
		fmt.Println("конец")
	*/

	var a int
	a = rand.Intn(time.Now().Nanosecond())
	fmt.Println(a, "полное число")

	var b = random(4)
	fmt.Println(b, "b переменная с 4 цыфрами")

}

// функция возвращает рандомное число
func random(s int) int {
	var a int
	var res int
	a = rand.Intn(time.Now().Nanosecond())
	switch s {
	case 1:
		res = a / 100000000

	case 2:
		res = a / 10000000

	case 3:
		res = a / 1000000

	case 4:
		res = a / 100000

	}
	return res
}
