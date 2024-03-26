package hard

import (
	"leets/s2/arrays"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFirstMissingPositive(t *testing.T) {
	assert.Equal(t, 3, firstMissingPositive0(arrays.Of(1, 2, 0)))
	assert.Equal(t, 2, firstMissingPositive0(arrays.Of(3, 4, -1, 1)))
	assert.Equal(t, 1, firstMissingPositive0(arrays.Of(7, 8, 9, 11, 12)))
	assert.Equal(t, 1, firstMissingPositive0(nil))
	assert.Equal(t, 1, firstMissingPositive0([]int{}))
}

func BenchmarkFirstMissingPositive0(t *testing.B) {
	assert.Equal(t, 3, firstMissingPositive0(arrays.Of(1, 2, 0)))
	assert.Equal(t, 2, firstMissingPositive0(arrays.Of(3, 4, -1, 1)))
	assert.Equal(t, 1, firstMissingPositive0(arrays.Of(7, 8, 9, 11, 12)))
	assert.Equal(t, 1, firstMissingPositive0(nil))
	assert.Equal(t, 1, firstMissingPositive0([]int{}))
}

func BenchmarkFirstMissingPositive1(t *testing.B) {
	assert.Equal(t, 3, firstMissingPositive1(arrays.Of(1, 2, 0)))
	assert.Equal(t, 2, firstMissingPositive1(arrays.Of(3, 4, -1, 1)))
	assert.Equal(t, 1, firstMissingPositive1(arrays.Of(7, 8, 9, 11, 12)))
	assert.Equal(t, 1, firstMissingPositive1(nil))
	assert.Equal(t, 1, firstMissingPositive1([]int{}))
}

func BenchmarkFirstMissingPositive(t *testing.B) {
	assert.Equal(t, 3, firstMissingPositive(arrays.Of(1, 2, 0)))
	assert.Equal(t, 2, firstMissingPositive(arrays.Of(3, 4, -1, 1)))
	assert.Equal(t, 1, firstMissingPositive(arrays.Of(7, 8, 9, 11, 12)))
	assert.Equal(t, 1, firstMissingPositive(nil))
	assert.Equal(t, 1, firstMissingPositive([]int{}))
}
