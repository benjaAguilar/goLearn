package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Should repeat A 5 times", func(t *testing.T) {
		got := Repeat("A", 5)
		expect := "AAAAA"

		if got != expect {
			t.Errorf("- got: %v \n - expect: %v", got, expect)
		}
	})

	t.Run("Can specify the number of repetitions", func(t *testing.T) {
		got := Repeat("f", 3)
		expect := "fff"

		if got != expect {
			t.Errorf("- got: %v \n - expect: %v", got, expect)
		}
	})

	t.Run("0 returns empty string", func(t *testing.T) {
		got := Repeat("f", 0)
		expect := ""

		if got != expect {
			t.Errorf("- got: %v \n - expect: %v", got, expect)
		}
	})

	t.Run("negative nums repeat the same times", func(t *testing.T) {
		got := Repeat("f", -2)
		expect := "ff"

		if got != expect {
			t.Errorf("- got: %v \n - expect: %v", got, expect)
		}
	})
}

func Benchmark(b *testing.B) {
	for b.Loop() {
		Repeat("A", 5)
	}
}

func ExampleRepeat() {
	out := Repeat("t", 4)
	fmt.Println(out)
	// Output:
	// tttt
}
