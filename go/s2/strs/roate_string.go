package strs

func rotateString(s string, goal string) bool {

	l := len(s)
	if l != len(goal) {
		return false
	}

	for i := 0; i < l; i++ {
		mtch := true
		for j := 0; j < l; j++ {
			x := i + j
			if x >= l {
				x = x - l
			}

			if s[x] != goal[j] {
				mtch = false
				break
			}
		}

		if mtch {
			return true
		}
	}
	return false
}
