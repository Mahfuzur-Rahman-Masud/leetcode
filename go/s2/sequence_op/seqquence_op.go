package sequenceop

import (
	"strconv"
)

func calPoints(op []string) int {
	scr := make([]int, len(op))
	mark := 0

	for _, o := range op {
		if o == "+" {
			scr[mark] = scr[mark-1] + scr[mark-2]
		} else if o == "D" {
			scr[mark] = scr[mark-1] * 2
		} else if o == "C" {
			scr[mark-1] = 0
			mark -= 2
		} else {
			i, _ := strconv.Atoi(o)
			scr[mark] = i
		}

		mark++
	}

	out := 0
	for _, s := range scr[:mark] {
		out += s
	}

	return out
}
