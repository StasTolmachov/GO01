package main

import (
	"fmt"
	"testing"
)

type ValueTest struct {
	name string
	in   int
	out  int
}

func TestFib(t *testing.T) {
	ValueTests := []ValueTest{
		{"case 9", 9, 34},
		{"case 15", 15, 610},
		{"case 22", 22, 17711},
	}

	for _, tmp := range ValueTests {
		tmp := tmp
		t.Run(tmp.name, func(t *testing.T) {
			result := Fib(tmp.in)
			if result != tmp.out {
				t.Error("TestFib", result)
			}
		})

	}

}

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := Fib(15)
		_ = result
	}
}


func ExampleFib() {
	result := Fib(15)
	fmt.Println(result)
	// Output: 610
}

func TestFib2(t *testing.T) {
	ValueTests := []ValueTest{
		{"case 9", 9, 34},
		{"case 15", 15, 610},
		{"case 22", 22, 17711},
	}

	for _, tmp := range ValueTests {
		tmp := tmp
		t.Run(tmp.name, func(t *testing.T) {
			result := Fib2(tmp.in)
			if result != tmp.out {
				t.Error("TestFib", result)
			}
		})
	}

}

func BenchmarkFib2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := Fib2(15)
		_ = result
	}
}

func ExampleFib2() {
	result := Fib2(15)
	fmt.Println(result)
	// Output: 610
}