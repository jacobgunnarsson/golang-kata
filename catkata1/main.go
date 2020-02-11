package kata

import (
	"math"
)

func solve(yard []string, minDistance int) bool {
	cats := make(map[rune][2]int)

	for i, r := range yard {
		for j, c := range r {
			if string(c) != "-" {
				cats[c] = [2]int{i, j}
			}
		}
	}

	for k1, v1 := range cats {
		for k2, v2 := range cats {
			if k1 != k2 {
				distance := math.Sqrt(math.Pow(float64(v2[0] - v1[0]), 2) + math.Pow(float64(v2[1] - v1[1]), 2))

				if distance < float64(minDistance) {
					return false
				}
			}
		}
	}

	return true
}
