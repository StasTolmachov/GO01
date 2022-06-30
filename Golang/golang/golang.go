// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// func main() {
// 	var count = 0
// 	var num int

// 	for count < 10 { // Начало области видимости
// 		num = rand.Intn(10) + 1
// 		fmt.Println(num)

// 		count++
// 	} // Конец области видимости
// }

// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// var era = "AD" // переменная era доступна через пакет

// func main() {
// 	year := 2018 // переменные era и year находятся в области видимости

// 	switch month := rand.Intn(12) + 1; month { // переменные era, year и month в области видимости
// 	case 2:
// 		day := rand.Intn(28) + 1 // новый день
// 		fmt.Println(era, year, month, day)
// 	case 4, 6, 9, 11:
// 		day := rand.Intn(30) + 1
// 		fmt.Println(era, year, month, day)
// 	default:
// 		day := rand.Intn(31) + 1
// 		fmt.Println(era, year, month, day)
// 	} // month и day за пределами области видимости
// } // year за пределами области видимости

// package main

// import (
// 	"fmt"
// 	"math/rand"
// )

// var era = "AD"

// func main() {
// 	year := 2018
// 	month := rand.Intn(12) + 1
// 	daysInMonth := 31

// 	switch month {
// 	case 2:
// 		daysInMonth = 28
// 	case 4, 6, 9, 11:
// 		daysInMonth = 30
// 	}

// 	day := rand.Intn(daysInMonth) + 1
// 	fmt.Println(era, year, month, day)
// }

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var year int
	var month int
	var dayToday int
	//var daysInMonth int

	year = 2004
	//year = rand.Intn(2022)

	month = 9
	//month = rand.Intn(12) + 1
	//dayToday = 29
	//daysInMonth = 29

	for i := 0; i <= 10; i++ {
		if year%4 == 0 {
			switch month {
			case 2:
				dayToday = rand.Intn(29) + 1
				fmt.Println("высокосный февраль и 29 дней")

			case 4, 6, 9, 11:
				dayToday = rand.Intn(30) + 1
				fmt.Println("высокосный год в месяце до 30 дней")

			default:
				dayToday = rand.Intn(31) + 1
				fmt.Println("высокосный год в месяце до 31 дней")
			}
		} else {
			switch month {
			case 2:
				dayToday = rand.Intn(28) + 1
				fmt.Println("не высокосный февраль и 28 дней")

			case 4, 6, 9, 11:
				dayToday = rand.Intn(30) + 1
				fmt.Println("не высокосный год в месяце до 30 дней")

			default:
				dayToday = rand.Intn(31) + 1
				fmt.Println("не высокосный год в месяце до 31 дней")
			}

		}
		fmt.Println(year, month, dayToday)
	}
}
