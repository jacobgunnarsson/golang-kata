package kata

func solve(s string) (sum int) {
	for _, r := range s {
		sum += int(r) - 96
	}

	return sum
}
