package kata

func solve(numbers []int) int {
	sum := 0

	for _, v := range numbers {
		if v > 0 {
			sum += v
		}
	}

	return sum
}
