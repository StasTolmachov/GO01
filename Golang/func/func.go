package main

import "fmt"

func main() {
	number1 := 21
	number2 := 32
	a, b, c, d := calc(number1, number2)
	fmt.Println(a, b, c, d)

	number3 := 55
	number4 := 34
	a, b, c, d = calc(number3, number4)

	fmt.Println(a, b, c, d)

}

func calc(num1, num2 int) (sum, minus, multiply, division int) {
	sum = num1 + num2
	minus = num1 - num2
	multiply = num1 * num2
	division = num1 / num2
	return
}
