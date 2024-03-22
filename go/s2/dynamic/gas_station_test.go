package dynamic

import (
	"leets/s2/arrays"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGasStation(t *testing.T) {
	r := canCompleteCircuit(arrays.Of(1, 2, 3, 4, 5), arrays.Of(3, 4, 4, 1, 2))
	assert.Equal(t, 3, r)

	r = canCompleteCircuit(arrays.Of(2, 3, 4), arrays.Of(3, 4, 3))
	assert.Equal(t, -1, r)
}
