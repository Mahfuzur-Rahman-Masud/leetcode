package leet

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	return back_track(n, 0, "")
}

func back_track(left int, right int, s string) []string {
	if left == 0 && right == 0 {
		return []string{s}
	}

	out := []string{}

	if left > 0 {
		out = append(out, back_track(left-1, right+1, s+"(")...)
	}

	if right > 0 {
		out = append(out, back_track(left, right-1, s+")")...)
	}

	return out
}
