/*
package main

import "fmt"

func main() {
	const size = 3
	var arr = [size]int{3, 4, 5}

	var arr2 = arr

	fmt.Println(arr2)
}

*/
/*
package main

import "fmt"

func main() {
	var sl []int = []int{1, 2, 3}

	var sl2 []int = make([]int, 5, 10)

	fmt.Println(sl)
	fmt.Println(sl2)
}

*/

/*
package main

import "fmt"

func main() {
	var mp = map[string]int{"one": 1, "two": 2, "three": 3}

	mp["two"] = 5

	fmt.Println(mp["two"])
}
*/

/*
package main

import "fmt"

func main() {

	arr := [6]int{1, 2, 3, 4, 5}

	fmt.Println(arr)

}
*/
/*
package main

import "fmt"

func main() {
	arr := [...]int{5, 6, 7}

//	for i, el := range arr {
//		fmt.Println(i, el)
//	}

	for i := range arr {
		fmt.Println(i, arr[i])
	}
}
*/
/*
package main

import "fmt"

func main() {

	arr := []int{7, 4, 8, 9, 5}

	for i := range arr {

		p := arr[i] + 10
		fmt.Println(i, p)

	}

}
*/
/*
package main

import "fmt"

func main() {
	arr := [5][3]int{}
	arr[2][1] = 5
	a := arr
	fmt.Println(a)
}
*/
/*
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	inputNums := [5]int64{}
	scanner := bufio.NewScanner(os.Stdin)
	// записываем введённые числа в массив в обратном порядке
	for i := len(inputNums) - 1; i >= 0; i-- {
		if scanner.Scan() {
			num, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				panic(err)
			}
			inputNums[i] = num
		} else {
			panic("you must input 5 numbers")
		}
	}
	for _, el := range inputNums {
		fmt.Println(el)
	}
}
*/
/*
package main

import "fmt"

func main() {
	arr := [...]int{1, 2, 10, 4, 5, 6, 7, 8, 9, 10}
	sliceFromArr := arr[:5]
	slice := []int{1, 2, 3}
	fmt.Println("arr len:", len(arr))
	println()
	fmt.Println("sliceFromArr len:", len(sliceFromArr))
	fmt.Println("sliceFromArr cap:", cap(sliceFromArr))
	println()
	fmt.Println("slice len:", len(slice))
	fmt.Println("slice cap:", cap(slice))
	fmt.Println("значения слайса  ", sliceFromArr)
	fmt.Println(arr[2])
	fmt.Println(slice)
}


*/
/*
package main

import "fmt"

func main() {
	count := 5
	// 0 - начальная длина
	// count - начальная capacity
	slice := make([]int, 0, count)
	for i := 0; i < count; i++ {
		slice = append(slice, i)
		fmt.Println("cap:", cap(slice))
	}
	fmt.Println(slice)
}
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	arr := []int{}
	sc := bufio.NewScanner(os.Stdin)

	for sc.Scan() {
		num, err := strconv.ParseInt(sc.Text(), 10, 64)
		if err != nil {
			break
		}
		arr = append(arr, int(num))

	}

	InsertionSort(arr)
	fmt.Println(arr)

}

func InsertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		x := arr[i]
		j := i
		for ; j >= 1 && arr[j-1] > x; j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = x
	}
}
