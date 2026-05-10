package sum

import "testing"

func TestSum(t *testing.T) {
	t.Run("should sum all 5 the nums of an arr", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		expect := 15

		if got != expect {
			t.Errorf("\n - got: %v \n - expect: %v", got, expect)
		}
	})

	t.Run("can receive diferent arr sizes", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		expect := 6

		if got != expect {
			t.Errorf("\n - got: %v \n - expect: %v", got, expect)
		}
	})
}
