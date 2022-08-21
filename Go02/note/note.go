package main

import "fmt"

func culc(a, b int) int {
	res := a / b
	return res
}

func main() {
	var a int
	var b int
	// var res int

	a = 9
	b = 3

	defer fmt.Println(culc(a, b))
	fmt.Println("second")

}
