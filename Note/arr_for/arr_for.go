package main

import "fmt"

func main() {
	arr := []int{1, 2, 3, 4, 5}
	simplefor(arr)
}

func simplefor(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i])
	}
}

//todo
