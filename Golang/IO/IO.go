package main

import (
	"fmt"
	"os"
)

func main() {
	// fileText, err := os.OpenFile("/Users/st/go/src/GO01/Golang/IO/text.txt", os.O_APPEND, 0777)
	// if err != nil {
	// 	panic(err)
	// }
	// defer fileText.Close()

	// buf := make([]byte, 12)
	// fileText.Read(buf)

	// fmt.Println(buf, string(buf))

	fileText, err := os.ReadFile("/Users/st/go/src/GO01/Golang/IO/text.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fileText))

	os.WriteFile("/Users/st/go/src/GO01/Golang/IO/text.txt", []byte("goodbye"), 0777)

	

	fileText, err = os.ReadFile("/Users/st/go/src/GO01/Golang/IO/text.txt")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(fileText))
}
