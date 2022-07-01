// package main

// import "fmt"

// func main() {
// 	// var age int
// 	// var year int

// 	// fmt.Printf("любой текст%20v %4v\n", "SpaceX", 94)
// 	// fmt.Printf("любой текст%20v %4v\n", "Virgin Galactic", 100)

// 	// var a int
// 	// a -= 5
// 	// fmt.Println(a)

// }

package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num = rand.Intn(10) + 1
	fmt.Println(num)

	num = rand.Intn(100)
	fmt.Println(num)
}
