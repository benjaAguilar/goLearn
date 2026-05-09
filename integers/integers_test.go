package integers

import (
	"fmt"
	"testing"
)

func assertCorrectMessage(t testing.TB, got, expect int) {
	t.Helper()
	if got != expect {
		t.Errorf(
			"- got: %v \n - expected: %v",
			got, expect,
		)
	}
}

func TestAddFunction(t *testing.T) {
	t.Run("Should sum two numbers", func(t *testing.T) {
		got := Add(3, 3)
		expect := 6
		assertCorrectMessage(t, got, expect)
	})
}

func ExampleAdd() {
	sum := Add(3, 9)
	fmt.Println(sum)

	// Output:
	// 12
}
