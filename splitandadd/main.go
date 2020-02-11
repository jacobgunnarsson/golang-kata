package kata

import (
	"math"
)

func solve(numbers []int, n int) []int {
	if n == 0 || len(numbers) == 1 {
		return numbers
	}

	splitAt := int(math.Floor(float64(len(numbers) / 2)))
	left := numbers[:splitAt]
	right := numbers[splitAt:]

	for i := len(left) - 1; i >= 0; i-- {
		right[len(right) - 1 - (len(left) - 1 - i)] += left[i]
	}

	return solve(right, n - 1)
}
