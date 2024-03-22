package matrix

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxCount(t *testing.T) {
	assert.Equal(t, 4, maxCount(3, 3, [][]int{{2, 2}, {3, 3}}))
	assert.Equal(t, 9, maxCount(3, 3, [][]int{}))
	assert.Equal(t, 4, maxCount(3, 3, [][]int{{2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}, {2, 2}, {3, 3}, {3, 3}, {3, 3}}))

}
