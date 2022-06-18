package main

import (
	"fmt"
	"math"
)

func main() {

	var s, d, c float64

	fmt.Print("Введите площадь: ")
	fmt.Scanln(&s)

	d = (math.Sqrt((s / math.Pi))) * 2

	c = d * math.Pi

	fmt.Println("Диаметр: ", d, "\nДлина окружности: ", c)
}
