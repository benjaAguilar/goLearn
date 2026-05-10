package sum

func Sum(arr []int) int {
	var sum int

	for _, val := range arr {
		sum += val
	}

	return sum
}

func SumAll(numsToSum ...[]int) []int {
	lengthOfNums := len(numsToSum)
	sums := make([]int, lengthOfNums)

	for i, val := range numsToSum {
		sums[i] = Sum(val)
	}

	return sums
}
