package backtracking

import (
	"fmt"
	"leets/s2/arrays"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubsetXORSum(t *testing.T) {
	// assert.Equal(t, 6, subsetXORSum(arrays.Of(1, 3)))
	assert.Equal(t, 28, subsetXORSum(arrays.Of(5, 1, 6)))
}

func TestXor(t *testing.T) {
	fmt.Println(1 ^ 1)
	fmt.Println(1 ^ 0)
	fmt.Println(0 ^ 0)
	fmt.Println(0 ^ 1)

	fmt.Println(^1)
	x := [5]int{}
	y := x[5:]
	fmt.Println(y)

}
