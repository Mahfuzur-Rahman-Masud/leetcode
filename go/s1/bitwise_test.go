package s1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountBits(t *testing.T) {
	assert.Equal(t, []int{0, 1, 1}, countBits(2))
	assert.Equal(t, []int{0, 1, 1, 2, 1, 2}, countBits(5))
}

func TestDivideBy2(t *testing.T) {
	assert.Equal(t, 3, divideBy2(6))
	assert.Equal(t, 2, divideBy2(4))
	assert.Equal(t, 3, divideBy2(7))
	assert.Equal(t, 1, divideBy2(2))
}

func TestSetBits(t *testing.T) {
	assert.Equal(t, 2, countSetBits(5))
}

func TestReverseBit(t *testing.T) {
	m := map[uint32]uint32{43261596: 964176192,
		4294967293: 3221225471,
	}

	for i, o := range m {
		fmt.Println(i, o)
		r := reverseBits(i)
		assert.Equal(t, o, r)
	}
}

func TestHammingWeight(t *testing.T) {
	m := map[uint32]int{3: 2,
		0:  0,
		1:  1,
		8:  1,
		15: 4,
		5:  2,
	}

	for i, o := range m {
		fmt.Println(i, o)
		r := hammingWeight(i)
		assert.Equal(t, o, r)
	}
}
