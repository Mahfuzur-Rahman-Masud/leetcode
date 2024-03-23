package listnode

import (
	"leets/s2/arrays"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReorder(t *testing.T) {
	l := ListOf(1, 2, 3, 4)
	reorderList(l)
	assert.Equal(t, arrays.Of(1, 4, 2, 3), l.ToSlice())

	l = ListOf(1, 2, 3, 4, 5)
	reorderList(l)
	assert.Equal(t, arrays.Of(1, 5, 2, 4, 3), l.ToSlice())
}
