package fibonachi

import "testing"

type ValueTest struct {
	in  int
	out int
}

func TestFib(t *testing.T) {
	ValueTests := []ValueTest{
		{9, 34},
		{15, 610},
		{22, 17711},
	}

	for _, tmp := range ValueTests {

		t.Run("TestFib", func(t *testing.T) {
			t.Parallel()
			result := Fib(tmp.in)

			if result != tmp.out {
				t.Error("TestFib", result)
			}
		})

	}

}

func TestFib2(t *testing.T) {
	ValueTests := []ValueTest{
		{9, 34},
		{15, 610},
		{22, 17711},
	}

	for _, tmp := range ValueTests {
		t.Run("TestFib2", func(t *testing.T) {
			t.Parallel()
			result := Fib2(tmp.in)

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

func BenchmarkFib2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		result := Fib2(15)
		_ = result
	}
}
