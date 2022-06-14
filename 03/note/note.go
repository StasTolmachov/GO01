/*
package main

import "fmt"

func main() {
	a := 103
	if a%10 == 0 {

		fmt.Println("Число а кратно десяти")
	}
}
*/

/*
package main

import "fmt"

func main() {
	a := 33
	if a%10 == 0 {
		fmt.Println("Число а кратно десяти")

	} else if a%5 == 0 {
		fmt.Println("Число a не кратно десяти, но кратно пяти")
	} else if a%3 == 0 {
		fmt.Println("Число a не кратно десяти, но кратно 3")
	} else {
		fmt.Println("Число а не кратно десяти")
	}

}
*/



/*
package main

import "fmt"

func main() {
	a := 33
	if a%10 == 0 {
		fmt.Println("Число а кратно десяти")

	} else if a%5 == 0 {
		fmt.Println("Число a не кратно десяти, но кратно пяти")
	} else if a%3 == 0 {
		fmt.Println("Число a не кратно десяти, но кратно 3")
	} else {
		fmt.Println("Число а не кратно десяти")
	}

}
*/

/*
package main

import (
	"fmt"
)

func main() {
	s := "Hello, world!"
	for i := 0; i < len(s); i++ {
		fmt.Println(string(s[i]))
	}
}
*/

/*
package main

import (
	"fmt"
	"math/rand"
)

func main() {
	condition := true
	for condition {
		nr := rand.Int()
		fmt.Println(nr)
		if nr%4 == 0 {
			condition = false
		}
	}
}
*/

/*
package main

import (
	"fmt"
)

func main() {
	a := 33
	switch a % 10 {
	case 0:
		fmt.Println("Число a кратно десяти")
	case 5:
		fmt.Println("Число a не кратно десяти, но кратно пяти")
	case 3:
		fmt.Println("Число a кратно 3")
	default:
		fmt.Println("Число a не кратно ни десяти, ни пяти")
	}
}
*/







// calculator
package main

import (
	"fmt"
	"os"
	"math"
)

    var a, b, res float64
	var op string


func factorial(num float64) float64{
   if num == 1 || num == 0{
      return num
   }
   return num * factorial(num-1)
}

func main() {
	

	fmt.Print("Enter first number A: ")
	fmt.Scan(&a)
	
	fmt.Print("Enter an arithmetic operation(+ - * / expo fact): ")
	fmt.Scan(&op)

	if op == "expo" {
		var expo float64
		fmt.Print("Enter exponentiation: ")
	    fmt.Scan(&expo)
		res = math.Pow(float64(a), expo)
		fmt.Printf("The result of the operation: %.2f", res)

	} else if op == "fact" {
		res = factorial(a)
		fmt.Printf("The result of the operation: %.2f", res)

	} else {
	fmt.Print("Enter second number B: ")
	fmt.Scan(&b)


if b == 0 && op == "/" {
	fmt.Println("Error")
  } else {

	switch op {
	case "+":
		res = a + b

	case "-":
		res = a - b

	case "*":
		res = a * b

	case "/":
		res = a / b

	default:
		fmt.Println("Operation selected incorrectly")
		os.Exit(1)

	}

	fmt.Printf("The result of the operation: %.2f", res)
  }
}
}



/*

package main
import "fmt"
func factorial(num int) int{
   if num == 1 || num == 0{
      return num
   }
   return num * factorial(num-1)
}
func main(){
   fmt.Println(factorial(3))
   fmt.Println(factorial(4))
   fmt.Println(factorial(5))
}



/*
package main

import (
	"fmt"
)

func main() {
	fmt.Println(bubbleSort([]int{5, 3, 6, 8, 1, 2}))
}
func bubbleSort(arr []int) []int {
	swapped := true
	for swapped {
		swapped = false
		for i := 0; i < len(arr)-1; i++ {
			if arr[i+1] < arr[i] {
				arr[i+1], arr[i] = arr[i], arr[i+1]
				swapped = true
			}
		}
	}
	return arr
}
*/