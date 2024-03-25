package algo

import (
	"leets/s2/arrays"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinarySearch(t *testing.T) {
	assert.Equal(t, 4, search(arrays.Of(-1, 0, 3, 5, 9, 12), 9))
	assert.Equal(t, -1, search(arrays.Of(-1, 0, 3, 5, 9, 12), 2))
}
