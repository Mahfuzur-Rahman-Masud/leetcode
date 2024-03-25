package s1

func MySqrt(x int) int {
	if x == 0 || x == 1 {
		return x
	}

	start := 1
	end := x
	var mid int

	for start <= end {
		mid = start + (end-start)/2

		sqr := mid * mid
		if sqr == x {
			return mid
		}

		if sqr > x {
			end = mid - 1
		} else {
			start = mid + 1
		}
	}

	return end
}
