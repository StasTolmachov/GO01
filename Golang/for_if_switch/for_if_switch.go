// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
// 	var comand = "красное черное"
// 	var exit = strings.Contains(comand, "красное")
// 	fmt.Println(exit)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	fmt.Println("На дворе 2100 год. Он високосный?")

// 	var year = 2104
// 	var leap = year%400 == 0 || (year%4 == 0 && year%100 != 0)

// 	if leap {
// 		fmt.Println("Этот год високосный!")
// 	} else {
// 		fmt.Println("К сожалению, нет. Этот год не високосный.")
// 	}

// 	fmt.Println(leap)
// }

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	var a = 6

// 	var b = a == 5 || a%3 == 0

// 	fmt.Println(!b)

// }

// package main

// import "fmt"

// func main() {
// 	var a = 1
// 	for a <= 10 {
// 		fmt.Println(a)
// 		a++
// 	}
// }

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var numUser int
	var numComp int

	numUser = 5

	for numUser != numComp {
		numComp = rand.Intn(10) + 1
		fmt.Println(numComp)
	}

}
