package kata

import (
	"sort"
)

func solve(arr []int) []int {
	sort.Sort(sort.Reverse(sort.IntSlice(arr)))

	valley := make([]int, len(arr))
	leftIndex := 0
	rightIndex := len(valley) - 1

	for i, v := range arr {
		if i%2 == 0 {
			valley[leftIndex] = v
			leftIndex++
		} else {
			valley[rightIndex] = v
			rightIndex--
		}
	}

	return valley
}
