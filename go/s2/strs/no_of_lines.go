package strs

func numberOfLines(widths []int, s string) []int {
	l, w := 1, 0

	x := 0
	for _, c := range s {
		c = c - 97
		xx := widths[c]

		if x+xx > 100 {
			w = x
			l++
			x = xx
		} else if x+xx == 100 {
			w = 100
			l++
			x = 0
		} else {
			x += xx
		}
	}

	if x > 0 {
		w = x
	} else {
		l--
	}

	return []int{l, w}
}
