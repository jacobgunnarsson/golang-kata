package kata

import (
	"sort"
	"strings"
)

func solve(s string) string {
	people := strings.Split(s, ";")

	for i, p := range people {
		name := strings.Split(strings.ToUpper(p), ":")

		people[i] = strings.Join([]string{name[1], name[0]}, ", ")
	}

	sort.Strings(people)

	return "(" + strings.Join(people, ")(") + ")"
}
