package s1

import (
	"fmt"
	"testing"
)

func TestMySqrt(t *testing.T) {
	m := map[int]int{
		4:  2,
		8:  2,
		1:  1,
		0:  0,
		11: 3,
		10: 3,
	}

	for k, v := range m {
		r := MySqrt((k))
		fmt.Printf("in: %d\t out: %d|%d\n", k, v, r)
		if r != v {
			t.Error("Assertion failed!")
		}
	}
}
