package binary

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAlternatingBits(t *testing.T) {
	// var x int32 = 0x55555555
	// var y int32 = 0x2aaaaaaa
	// fmt.Printf("%32b\n", x)
	// fmt.Printf("%32b\n", y)

	// fmt.Printf("%32b\n", 0x55555555)

	assert.True(t, hasAlternatingBits(5))
	assert.True(t, !hasAlternatingBits(7))
	assert.True(t, !hasAlternatingBits(11))
	assert.True(t, !hasAlternatingBits(4))
	assert.True(t, hasAlternatingBits(0))
	assert.True(t, hasAlternatingBits(01))

}
