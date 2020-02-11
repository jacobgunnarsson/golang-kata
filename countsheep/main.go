package kata

import "strconv"

func countsheep(num int) (sheep string) {
	for i := 1; i <= num; i++ {
		sheep += strconv.Itoa(i) + " sheep..."
	}

	return sheep
}
