package main

import "testing"

func TestFib(t *testing.T) {
	result := 33

	if result != fib(9) {
		t.Error(result)
	}
}
