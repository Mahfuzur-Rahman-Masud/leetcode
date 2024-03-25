package strs

func toLowerCase(s string) string {
	bs := []byte(s)

	for i, b := range bs {
		if b > 64 && b < 91 {
			bs[i] = b + 32
		}
	}

	return string(bs)
}
