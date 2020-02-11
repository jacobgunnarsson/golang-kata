package kata

import (
	"strings"
)

func solve(str string) (highest int) {
	const vowels = "aeiou"

	current := 0

	for _, r := range str {
		if !strings.ContainsRune(vowels, r) {
			current += int(r) - 96

			if (current > highest) {
				highest = current
			}
		} else {
			current = 0
		}
	}

	return highest
}
