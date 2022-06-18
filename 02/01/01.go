package main

import "fmt"

func main() {

	var a, b int

	fmt.Print("Введите длину: ")
	fmt.Scanln(&a)
	fmt.Print("Введите ширину: ")
	fmt.Scanln(&b)

	fmt.Println("Площадь:", a*b)
}
