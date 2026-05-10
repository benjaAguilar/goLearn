package sum

import (
	"reflect"
	"slices"
	"testing"
)

func assertCorrectMessage(t testing.TB, got, expect any) {
	t.Helper()
	if got != expect {
		t.Errorf(
			"\n - got: %v \n - expected: %v",
			got, expect,
		)
	}
}

func TestSum(t *testing.T) {
	t.Run("should sum all 5 the nums of an arr", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		expect := 15

		assertCorrectMessage(t, got, expect)
	})

	t.Run("can receive diferent arr sizes", func(t *testing.T) {
		numbers := []int{1, 2, 3}
		got := Sum(numbers)
		expect := 6

		assertCorrectMessage(t, got, expect)
	})
}

func TestSumAll(t *testing.T) {
	t.Run("should recieve 2 arr and return on with both sums", func(t *testing.T) {
		arr_one := []int{1, 2}
		arr_two := []int{9, 3, 1}

		got := SumAll(arr_one, arr_two)
		expect := []int{3, 13}

		if !slices.Equal(got, expect) {
			t.Errorf(
				"\n - got: %v \n - expected: %v",
				got, expect,
			)
		}
	})
}

func TestSumAllTails(t *testing.T) {
	t.Run("should sum all tails", func(t *testing.T) {
		got := SumAllTails([]int{1, 2, 3}, []int{4, 9, 3})
		expect := []int{5, 12}

		if !reflect.DeepEqual(got, expect) {
			t.Errorf(
				"\n - got: %v \n - expected: %v",
				got, expect,
			)
		}
	})

	t.Run("safely sum empty arrays", func(t *testing.T) {
		got := SumAllTails([]int{}, []int{4, 9, 3})
		expect := []int{0, 12}

		if !reflect.DeepEqual(got, expect) {
			t.Errorf(
				"\n - got: %v \n - expected: %v",
				got, expect,
			)
		}
	})
}
