package kata

import "strconv"

func solve(num int) (sheep string) {
	for i := 1; i <= num; i++ {
		sheep += strconv.Itoa(i) + " sheep..."
	}

	return sheep
}
