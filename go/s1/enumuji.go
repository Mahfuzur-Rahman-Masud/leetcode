package s1

import (
	"fmt"
)

type Color int

const (
	Red Color = iota
	Blue
	Green
)

func TakeColor(c Color) {
	switch c {
	case Red:
		fmt.Println("RED")
	case Green:
		fmt.Println("GREEN")
	case Blue:
		fmt.Println("BLUE")
	default:
		fmt.Println("NOT A Color")

	}
}
