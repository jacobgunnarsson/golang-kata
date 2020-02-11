package kata

import (
	"strings"
)

func solve(s string) (longest int) {
	const vowels = "aeiou"
	current := 0

	for _, r := range s {
		if strings.ContainsRune(vowels, r) {
			current++
		} else {
			if current > longest {
				longest = current
			}

			current = 0
		}
	}

  return longest
}
