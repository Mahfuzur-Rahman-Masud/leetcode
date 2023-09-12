package leet

func convertToTitle(col int) string {
	bytes := make([]byte, 7) // 26^7 over flows i32

	pos := 6 // last index of bytes
	for col > 0 {
		col--
		bytes[pos] = byte((col % 26) + 65) // 65 = ASCII 'A'
		pos--
		col /= 26
	}

	return string(bytes[pos+1:])
}

func titleToNumber(columnTitle string) int {
	b := []byte(columnTitle)

	l := len(b)
	out := 0

	for i := 0; i < l; i++ {
		idx := l - 1 - i
		// fmt.Println((b[idx])-64, i)
		out += pow(26, i) * int(b[idx]-64)
	}

	return out
}

func pow(a, b int) int {
	// fmt.Println("POW: ", a, b)
	if b == 0 {
		return 1
	}

	if a == 0 || a == 1 {
		return a
	}

	out := a
	for b > 1 {
		out *= a
		b--
	}

	return out
}
