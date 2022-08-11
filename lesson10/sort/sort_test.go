package sort

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSort(t *testing.T) {

	arr := []int{99, 34, 2, 65, 44}
	result := []int{2, 34, 44, 65, 99}

	InsertionSort(arr)
	if reflect.DeepEqual(arr, result) != true {
		t.Error("InsertionSort", arr, result)
	}

}

func ExampleSort() {
	arr := []int{99, 34, 2, 65, 44}
	InsertionSort(arr)
	fmt.Println(arr)
	// Output: 2, 34, 44, 65, 99
}
