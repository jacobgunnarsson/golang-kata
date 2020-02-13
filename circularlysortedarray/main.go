package kata

import (
	"sort"
)

func solve(r []int) bool {
	for i, v := range r {
		if i < len(r) - 1 && v > r[i + 1]  {
			return sort.IntsAreSorted(append(r[i + 1:], r[:i + 1]...))
		}
	}

	return true
}
