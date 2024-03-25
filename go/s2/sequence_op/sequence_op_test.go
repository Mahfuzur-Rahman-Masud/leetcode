package sequenceop

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSequenceOp(t *testing.T) {
	assert.Equal(t, 30, calPoints([]string{"5", "2", "C", "D", "+"}))
	assert.Equal(t, 27, calPoints([]string{"5", "-2", "4", "C", "D", "9", "+", "+"}))
	assert.Equal(t, 0, calPoints([]string{"1", "C"}))
}
