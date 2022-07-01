
// высокосный год
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var year int
	var month int
	var dayToday int
	var year4 bool

	year = 2000 + rand.Intn(20)

	month = rand.Intn(12) + 1

	for i := 0; i <= 10; i++ {

		switch month {
		case 2:
			if year4 {
				dayToday = rand.Intn(29) + 1
			} else {
				dayToday = rand.Intn(28) + 1
			}

		case 4, 6, 9, 11:
			dayToday = rand.Intn(30) + 1

		default:
			dayToday = rand.Intn(31) + 1

		}
		fmt.Println(year, month, dayToday)
	}

}
