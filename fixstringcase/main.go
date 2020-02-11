package kata

import (
	"strings"
	"unicode"
)

func solve(str string) string {
	uppercase := 0

	for _, r := range str {
		if unicode.IsUpper(r) {
			uppercase++
		}
	}


	if uppercase > len(str) / 2 {
		return strings.ToUpper(str)
	}

	return strings.ToLower(str)
}
