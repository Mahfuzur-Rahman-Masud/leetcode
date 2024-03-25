package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindShortestSubArray(t *testing.T) {
	assert.Equal(t, 2, findShortestSubArray(Of(1, 2, 2, 3, 1)))
	assert.Equal(t, 6, findShortestSubArray(Of(1, 2, 2, 3, 1, 4, 2)))
	assert.Equal(t, 7, findShortestSubArray(Of(2, 1, 1, 2, 1, 3, 3, 3, 1, 3, 1, 3, 2)))
}
