package kata

import (
	"strings"
	"strconv"
)

func solve(s string) string {
	occurences := make(map[rune]int)
	numericals := make([]string, len(s))

	for _, r := range s {
		occurences[r]++
		numericals = append(numericals, strconv.Itoa(occurences[r]))
	}

	return strings.Join(numericals[:], "")
}
