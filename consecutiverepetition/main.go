package kata

type Result struct {
	C rune
	L int
}

func solve(text string) (result Result) {
	curLen := make(map[rune]int)

	for i, r := range text {
		curLen[r]++

		if i == len(text) - 1 || (i + 1 < len(text) && r != rune(text[i + 1])) {
			if (curLen[r] > result.L) {
				result.C = r
				result.L = curLen[r]
			}

			curLen[r] = 0
		}
	}

	return result
}
