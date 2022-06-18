package main

import "fmt"

func main() {

	var a int
	var s int
	var d int
	var e int

	fmt.Print("Введите трехзначное число: ")
	fmt.Scanln(&a)

	s = a / 100

	d = (a - (s * 100)) / 10
	e = a - (s * 100) - (d * 10)

	fmt.Println("Сотни: ", s, "Десятки: ", d, "Единицы: ", e)

}
