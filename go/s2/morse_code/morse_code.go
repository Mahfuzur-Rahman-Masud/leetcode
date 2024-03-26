package morsecode

var mcode = []string{".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}

func uniqueMorseRepresentations(words []string) int {
	container := map[string]struct{}{}
	for _, word := range words {
		t := ""
		for _, c := range word {
			i := c - 97
			t += mcode[i]
		}

		container[t] = struct{}{}
	}

	return len(container)
}
