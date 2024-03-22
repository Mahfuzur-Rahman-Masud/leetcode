package matrix

func maxCount(m int, n int, ops [][]int) int {
	for _, o := range ops {
		m2 := o[0]
		n2 := o[1]
		if m2 < m {
			m = m2
		}

		if n2 < n {
			n = n2
		}
	}

	return m * n
}
