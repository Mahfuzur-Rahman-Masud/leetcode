package listnode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverse(t *testing.T) {
	assert.Equal(t, SliceOf(5, 4, 3, 2, 1), reverseList(ListOf(1, 2, 3, 4, 5)).ToSlice())
	assert.Equal(t, SliceOf(2, 1), reverseList(ListOf(1, 2)).ToSlice())
	assert.Equal(t, SliceOf(1), reverseList(ListOf(1)).ToSlice())

}

func TestReverseValueSwitch(t *testing.T) {
	assert.Equal(t, SliceOf(5, 4, 3, 2, 1), reverseListValueSwitch(ListOf(1, 2, 3, 4, 5)).ToSlice())
	assert.Equal(t, SliceOf(2, 1), reverseListValueSwitch(ListOf(1, 2)).ToSlice())
	assert.Equal(t, SliceOf(1), reverseListValueSwitch(ListOf(1)).ToSlice())
}

func BenchmarkReverse(t *testing.B) {
	assert.Equal(t, SliceOf(5, 4, 3, 2, 1), reverseList(ListOf(1, 2, 3, 4, 5)).ToSlice())
	assert.Equal(t, SliceOf(2, 1), reverseList(ListOf(1, 2)).ToSlice())
	assert.Equal(t, SliceOf(1), reverseList(ListOf(1)).ToSlice())

}

func BenchmarkReverseValueSwitch(t *testing.B) {
	assert.Equal(t, SliceOf(5, 4, 3, 2, 1), reverseListValueSwitch(ListOf(1, 2, 3, 4, 5)).ToSlice())
	assert.Equal(t, SliceOf(2, 1), reverseListValueSwitch(ListOf(1, 2)).ToSlice())
	assert.Equal(t, SliceOf(1), reverseListValueSwitch(ListOf(1)).ToSlice())
}
