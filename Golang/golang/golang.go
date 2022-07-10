// package main

// import (
// 	"fmt"
// 	"unicode/utf8"
// )

// func main() {

// 	a := "massage"
// 	b := utf8.RuneCountInString(a)
// 	fmt.Println(b)

// 	for i := 0; i <= b; i++ {
// 		fmt.Printf("%c\n", a[i])
// 	}
// }

// package main

// import (
//     "fmt"
//     "unicode/utf8"
// )

// func main() {
//     question := "¿Cómo estás?"

//     fmt.Println(len(question), "bytes") // Выводит: 15 bytes
//     fmt.Println(utf8.RuneCountInString(question), "runes") // Выводит: 12 runes

//     c, size := utf8.DecodeRuneInString(question)
//     fmt.Printf("First rune: %c %v bytes", c, size) // Выводит: First rune: ¿ 2 bytes
// }

// package main

// import "fmt"

// type celsius float64
// type fahrenheit float64
// type kelvin float64

// // высчитывание цельсия
// func (f fahrenheit) celsius() celsius {
// 	return celsius((f - 32.0) * 5.0 / 9.0)
// }
// func (k kelvin) celsius() celsius {
// 	return celsius(k - 273.15)
// }

// // высчитывание кельвина
// func (c celsius) kelvin() kelvin {
// 	return kelvin(c + 273.15)
// }
// func (f fahrenheit) kelvin() kelvin {
// 	return f.celsius().kelvin()
// }

// // высчитывание фаренгейт
// func (c celsius) fahrenheit() fahrenheit {
// 	return fahrenheit((c * 9.0 / 5.0) + 32.0)
// }
// func (k kelvin) fahrenheit() fahrenheit {
// 	return k.celsius().fahrenheit()
// }

// func main() {
// 	var kel kelvin
// 	var c celsius
// 	kel = 294
// 	c = kel.celsius()
// 	fmt.Println(kel, c)
// }

// package main

// import "fmt"

// type uah float64
// type usd float64
// type zloty float64

// // uah
// func (us usd) uah() uah {
// 	return uah(us * 35.5)
// }
// func (zl zloty) uah() uah {
// 	return uah(zl * 6.28)
// }

// //usd
// func (ua uah) usd() usd {
// 	return usd(ua / 35.5)
// }
// func (zl zloty) usd() usd {
// 	return usd(zl * 0.21)
// }

// // zloty
// func (us usd) zloty() zloty {
// 	return zloty(us / 0.21)
// }
// func (ua uah) zloty() zloty {
// 	return zloty(ua / 6 / 28)
// }

// func main() {
// 	var zlotyBogdan zloty
// 	var us usd
// 	var ua uah

// 	zlotyBogdan = 4000

// 	us = zlotyBogdan.usd()
// 	ua = zlotyBogdan.uah()

// 	fmt.Println(us, ua)
// }
// package main

// import "fmt"

// func main() {
// 	planets := [...]string{ // Компилятор Go подсчитывает элементы
// 		"Меркурий",
// 		"Венера",
// 		"Земля",
// 		"Марс",
// 		"Юпитер",
// 		"Сатурн",
// 		"Уран",
// 		"Нептун", // Запятая в самом конце является обязательной
// 	}
// 	fmt.Println(len(planets))
// 	fmt.Println("-----------------")

// 	for i := 0; i < len(planets); i++ {
// 		planetsList := planets[i]
// 		fmt.Println(i, planetsList)
// 	}
// 	fmt.Println("-----------------")

// 	for i, plaplanetsList := range planets {
// 		fmt.Println(i, plaplanetsList)
// 	}
// 	fmt.Println("-----------------")

// 	for _, planetsList := range planets {
// 		fmt.Println(planetsList)
// 	}

// 	fmt.Println("-----------------")

// 	var board [8][8]string // Массив из восьми массивов с восемью строками

// 	board[0][0] = "r"
// 	board[0][7] = "r" // Ставит ладью на клетку с координатами [ряд][столбец]

// 	for column := range board[1] {
// 		board[1][column] = "p"
// 	}

// 	fmt.Println(board)
// }

// package main

// import "fmt"

// func display(board [8][8]rune) {
// 	for _, row := range board {
// 		for _, column := range row {
// 			if column == 0 {
// 				fmt.Print("  ")
// 			} else {
// 				fmt.Printf("%c ", column)
// 			}
// 		}
// 		fmt.Println()
// 	}
// }

// func main() {
// 	var board [8][8]rune

// 	// черные фигуры
// 	board[0][0] = 'r'
// 	board[0][1] = 'n'
// 	board[0][2] = 'b'
// 	board[0][3] = 'q'
// 	board[0][4] = 'k'
// 	board[0][5] = 'b'
// 	board[0][6] = 'n'
// 	board[0][7] = 'r'

// 	// пешки
// 	for column := range board[1] {
// 		board[1][column] = 'p'
// 		board[6][column] = 'P'
// 	}

// 	// белые фигуры
// 	board[7][0] = 'R'
// 	board[7][1] = 'N'
// 	board[7][2] = 'B'
// 	board[7][3] = 'Q'
// 	board[7][4] = 'K'
// 	board[7][5] = 'B'
// 	board[7][6] = 'N'
// 	board[7][7] = 'R'

// 	display(board)
// }

// package main

// import "fmt"

// func main() {

// 	planets := [...]string{ // Компилятор Go подсчитывает элементы
// 		"Меркурий",
// 		"Венера",
// 		"Земля",
// 		"Марс",
// 		"Юпитер",
// 		"Сатурн",
// 		"Уран",
// 		"Нептун", // Запятая в самом конце является обязательной
// 	}
// 	fmt.Println(planets)
// 	fmt.Println("-----------------")

// 	gasPlanets := planets[3:]
// 	fmt.Println(gasPlanets)

// 	fmt.Println("-----------------")

// 	planets[4] = "sdferf"

// 	fmt.Println(planets)
// 	fmt.Println(gasPlanets)

// 	fmt.Println("-----------------")

// 	fmt.Println(len(planets))
// 	fmt.Println(len(gasPlanets))

// 	fmt.Println("-----------------")

// 	gasPlanets = append(gasPlanets, "new")
// 	fmt.Println(planets)
// 	fmt.Println(gasPlanets)
// 	fmt.Println(len(planets))
// 	fmt.Println(len(gasPlanets))
// 	fmt.Println("---выделение срезов через make----")

// 	slice := make([]string, 0, 10)
// 	fmt.Println(len(slice))
// 	slice = append(slice, "sliceNew")
// 	fmt.Println(slice)
// 	fmt.Println(len(slice))

// }

package main

import "fmt"

func main() {
	planets := map[string]int{
		"first":  11,
		"second": 22,
		"three": 33,
	}
	fmt.Println(planets)

	moon := planets["first"]
	fmt.Println(moon)

	for _, p := range planets {
		
	}
}
