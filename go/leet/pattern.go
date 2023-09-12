package leet

func isMatch(s string, p string) bool {
	ls := len(s)
	lp := len(p)
	ip := 0
	x := byte('#')
	matchAny := false
	matchOnly := x
	dot := byte('.')
	more := byte('*')

	for i := 0; i < ls; i++ {
		c := s[i]

		if matchAny {
			continue
		}

		if c == matchOnly {
			continue
		}

		matchOnly = x

		if ip == lp {
			return false
		}

		pp := p[ip]
		ip++

		if pp == c {
			continue
		}

		if pp == dot {
			continue
		}

		if pp == more {
			if ip < 2 {
				return false
			}

			prev := p[ip-2]

			if prev == more {
				return false
			}

			if prev == dot {
				if ip == lp {
					return true
				}

				matchAny = true
				continue
			} else {
				matchOnly = prev
				if matchOnly != c {
					return false
				}
			}
		}
	}

	if ip == lp {
		return true
	}

	if ip == lp-1 && p[ip] == more {
		return p[ip-1] != more
	}

	if matchOnly == s[ls-1] {

	}

	return true

}
