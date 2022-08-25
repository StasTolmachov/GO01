package main

import (
	"lesson10/fibonachi"
	"lesson10/sort"
	"fmt"
)

func main() {

	res := fibonachi.Fib(15)
	fmt.Println(res)

	res2 := fibonachi.Fib2(15)
	fmt.Println(res2)

	arr := []int{99, 34, 2, 65, 44}
	fmt.Println(arr)
	sort.InsertionSort(arr)
	fmt.Println(arr)

}
