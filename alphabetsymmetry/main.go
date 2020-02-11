package kata

import (
	"strings"
)

func solve(slice []string) []int {
	result := make([]int, len(slice))

	for i, v := range slice {
		for j, r := range strings.ToLower(v) {
			if j == int(r) - 97 {
				result[i] += 1
			}
		}
	}

	return result
}
