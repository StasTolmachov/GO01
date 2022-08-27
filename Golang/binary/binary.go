package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30}
	var userNumber int
	fmt.Println("Введите число от 1 до 30")
	fmt.Scan(&userNumber)
	var sliceNumber int
	low := 0
	mid := len(slice) / 2
	high := len(slice)
	sliceNumber = slice[mid]

	for sliceNumber != userNumber {
		if sliceNumber < userNumber {
			low = mid
			mid = (low + high) / 2
			sliceNumber = slice[mid]
			fmt.Println("Каждая итерация вправо", sliceNumber)
		} else {
			high = mid
			mid = (low + high) / 2
			sliceNumber = slice[mid]
			fmt.Println("Каждая итерация влево", sliceNumber)
		}
	}
	fmt.Println("Результат", sliceNumber)

}
