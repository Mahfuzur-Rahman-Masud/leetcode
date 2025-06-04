package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	var m int
	var n int
	var n1 []int
	var n2 []int

	n1 = Of(1, 2, 3, 0, 0, 0)
	n2 = Of(2, 5, 6)
	m = 3
	n = 3
	merge(n1, m, n2, n)
	assert.Equal(t, Of(1, 2, 2, 3, 5, 6), n1)

	n1 = Of(1)
	n2 = Of()
	m = 1
	n = 0
	merge(n1, m, n2, n)
	assert.Equal(t, Of(1), n1)

	n1 = Of(4, 0, 0, 0, 0, 0)
	n2 = Of(1, 2, 3, 5, 6)
	m = 1
	n = 5
	merge(n1, m, n2, n)
	assert.Equal(t, Of(1, 2, 3, 4, 5, 6), n1)

}
