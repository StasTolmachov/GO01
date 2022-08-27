package main

import "fmt"

func main() {

	x := []float32{}
	var n float32

	fmt.Println("введите значение")
	fmt.Scan(&n)

	for i := n; i >= 1; {
		i = i / 2
		x = append(x, i)
	}
	res := len(x) - 1
	fmt.Println("значение логарифм ", res, x)


}
