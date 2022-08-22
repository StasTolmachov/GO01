package main

import (
	"fmt"
	"os"
	"runtime/debug"
)

func main() {
	Bufer, err := os.OpenFile("data.json", os.O_RDWR, 0777)
	if err != nil {
		defer fmt.Println(err)
		defer debug.PrintStack()
	}
	fmt.Println(Bufer)
}
