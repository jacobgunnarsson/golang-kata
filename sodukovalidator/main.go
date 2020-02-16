package kata

import (
	"math"
)

func IsUnique (slice []int) bool {
	uniqMap := make(map[int]struct{})

	for _, v := range slice {
		if _, ok := uniqMap[v]; ok {
			return false
		} else {
			uniqMap[v] = struct{}{}
		}
	}

	return true
}

func solve(m [][]int) bool {
	columns := [9][]int{}
	nonets := [9][]int{}

	for i, r := range m {
		if !IsUnique(r) {
			return false
		}

		ny := int(math.Floor(float64(i) / float64(3))) * 3

		for j, c := range r {
			nx := int(math.Floor(float64(j) / float64(3)))

			columns[i] = append(columns[i], c)
			nonets[nx + ny] = append(nonets[nx + ny], c)
		}

		if !IsUnique(columns[i]) {
			return false
		}
	}

	for _, n := range nonets {
		if !IsUnique(n) {
			return false
		}
	}

  return true
}
