package s1

var bad int
var seed int

func setSeed(s int) int {
	seed = s
	return s
}

func guess(num int) int {
	if num == seed {
		return 0
	} else if num > seed {
		return -1
	} else {
		return 1
	}

}

func guessNumber(n int) int {
	l, h := 0, n

	for l < h {
		m := l + (h-l)/2
		r := guess(m)
		if r == 0 {
			return m
		} else if r == 1 {
			l = m + 1
		} else {
			h = m - 1
		}
	}

	return l
}

func setBadVersion(firstBad int) {
	bad = firstBad
}

func isBadVersion(version int) bool {
	return version >= bad
}

func firstBadVersion(n int) int {
	low, high := 0, n

	for low <= high {
		mid := low + (high-low)/2

		if isBadVersion(mid) {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}

	return low
}
