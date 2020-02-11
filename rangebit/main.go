package kata

import (
	"strconv"
	"strings"
)

func solve(a, b int) (sum int) {
	for i := a; i <= b; i++ {
		sum += strings.Count(strconv.FormatInt(int64(i), 2), "1")
	}

	return sum
}
