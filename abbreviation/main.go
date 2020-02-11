package kata

import (
	"strings"
)

func abbreviation(name string) string {
	names := strings.Split(name, " ")

	for i, n := range names {
		names[i] = strings.ToUpper(n[:1])
	}

	return strings.Join(names, ".")
}
