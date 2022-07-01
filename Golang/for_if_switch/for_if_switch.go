/*
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
*/

package main

import (
	"fmt"
	"math/rand"
)

const secondsPerDay = 86400

func main() {
	distance := 62100000
	company := ""
	trip := ""

	fmt.Println("Spaceline        Days Trip type  Price")
	fmt.Println("======================================")

	for count := 0; count < 10; count++ {
		switch rand.Intn(3) {
		case 0:
			company = "Space Adventures"
		case 1:
			company = "SpaceX"
		case 2:
			company = "Virgin Galactic"
		}

		speed := rand.Intn(15) + 16                  // 16-30 km/s
		duration := distance / speed / secondsPerDay // days
		price := 20.0 + speed                        // millions

		if rand.Intn(2) == 1 {
			trip = "Round-trip"
			price = price * 2
		} else {
			trip = "One-way"
		}

		fmt.Printf("%-16v %4v %-10v $%4v\n", company, duration, trip, price)
	}
}
