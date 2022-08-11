package sort

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)
// Sort функция сортировки
func Sort() {
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
