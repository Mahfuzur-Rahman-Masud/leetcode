package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNumSubArrayProductsLessThanK(t *testing.T) {
	assert.Equal(t, 8, numSubarrayProductLessThanK(Of(10, 5, 2, 6), 100))
	assert.Equal(t, 0, numSubarrayProductLessThanK(Of(1, 2, 3), 0))
}
