// calculator
package main

import (
	"fmt"
	"math"
	"os"
)

var a, b, res float64
var op string

func factorial(num float64) float64 {
	if num == 1 || num == 0 {
		return num
	}
	return num * factorial(num-1)
}

func main() {

	fmt.Print("Enter first number A: ")
	fmt.Scan(&a)

	fmt.Print("Enter an arithmetic operation(+ - * / expo fact): ")
	fmt.Scan(&op)

	if op == "expo" {
		var expo float64
		fmt.Print("Enter exponentiation: ")
		fmt.Scan(&expo)
		res = math.Pow(float64(a), expo)
		fmt.Printf("The result of the operation: %.2f", res)

	} else if op == "fact" {
		res = factorial(a)
		fmt.Printf("The result of the operation: %.2f", res)

	} else {
		fmt.Print("Enter second number B: ")
		fmt.Scan(&b)

		if b == 0 && op == "/" {
			fmt.Println("Error")
		} else {

			switch op {
			case "+":
				res = a + b

			case "-":
				res = a - b

			case "*":
				res = a * b

			case "/":
				res = a / b

			default:
				fmt.Println("Operation selected incorrectly")
				os.Exit(1)

			}

			fmt.Printf("The result of the operation: %.2f", res)
		}
	}
}