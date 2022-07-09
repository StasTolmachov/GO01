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
package main

func main() {

}
