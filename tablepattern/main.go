package kata

import (
	"fmt"
	"strings"
)

func createRow(columns int, last bool) (row string) {
	row += strings.Repeat("+---", columns) + "+"

	if (!last) {
		return row + "\n"
	}

	return row
}

func solve(rows int, columns int, str string) (table string) {
	currentIndex := 0
	table += createRow(columns, false)

	for i := 0; i < rows; i++ {
		row := ""

		for j := 0; j < columns; j++ {
			if (currentIndex < len(str)) {
				row += fmt.Sprintf("| %s ", string(str[currentIndex]))
			} else {
				row += "|   "
			}

			currentIndex++
		}

		table += fmt.Sprintf("%s|\n", row)
		table += createRow(columns, i == rows - 1)
	}

	return table
}
