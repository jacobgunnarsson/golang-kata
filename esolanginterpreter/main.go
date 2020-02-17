package kata

func solve(code string) string {
	var mem byte = 0
	bytes := []byte{}

	for _, r := range code {
		switch r {
			case '+': mem++
			case '.': bytes = append(bytes, mem)
		}
	}

  return string(bytes)
}
