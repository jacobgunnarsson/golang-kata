package kata

import (
	"strings"
)

func solve(words string) []string {
	wave := make([]string, 0)

	for i, r := range words {
		if r == ' ' {
			continue
		}

		wave = append(wave, words[:i] + strings.ToUpper(string(r)) + words[i + 1:])
	}

	return wave
}
