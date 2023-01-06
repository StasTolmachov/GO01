package main

import "fmt"

// kelvinToCelsius конвертирует °K в °C
// func kelvinToCelsius(k float64) float64 { // Объявляет функцию, что принимает параметр и возвращает результат
// 	k -= 273.15
// 	return k
// }

// func main() {
// 	kelvin := 294.0
// 	celsius := kelvinToCelsius(kelvin)           // Вызывает функцию передачи kelvin как первого аргумента
// 	fmt.Print(kelvin, "° K is ", celsius, "° C") // Выводит: 294° K is 20.850000000000023° C
// }

type celsius float64
type kelvin float64

var cel celsius = 123
var kel kelvin

// функция
func celsiusToKelvin(c celsius) kelvin {
	var k kelvin
	k = kelvin(c + 273)
	return (k)
}

// метод
func (c celsius) celsiusToKelvin() kelvin {
	var k kelvin
	k = kelvin(c + 273)
	return (k)
}

func main() {

	fmt.Println(celsiusToKelvin(123))
	fmt.Println(cel.celsiusToKelvin())
}
