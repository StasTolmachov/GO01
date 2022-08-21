package main

import (
	"fmt"
)

func culc(a, b int) (int, error) {
	var err error
	var res int
	res = a / b
	return res, err
}

type ResultStruct struct {
	Res int
	Err error
}

func main() {
	var a int
	var b int

	var result ResultStruct

	a = 9
	b = 0
	
	result.Res, result.Err = culc(a, b)

	fmt.Println(result)

	fmt.Println("second")

}
