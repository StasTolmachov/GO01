package main

import "fmt"

func main() {

	var a string

	fmt.Print("Введите трехзначное число: ")
	fmt.Scanln(&a)

	fmt.Println("Сотни: ", string(a[0]), "\nДесятки: ", string(a[1]), "\nЕдиницы: ", string(a[2]))

}
