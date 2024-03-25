package s1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGuessNumber(t *testing.T) {
	assert.Equal(t, setSeed(6), guessNumber(10))
	assert.Equal(t, setSeed(0), guessNumber(10))
	assert.Equal(t, setSeed(9), guessNumber(10))
	assert.Equal(t, setSeed(10), guessNumber(10))
	assert.Equal(t, setSeed(5), guessNumber(10))
	assert.Equal(t, setSeed(7), guessNumber(10))
	assert.Equal(t, setSeed(7), guessNumber(11))
	assert.Equal(t, setSeed(6), guessNumber(11))
	assert.Equal(t, setSeed(5), guessNumber(11))
	assert.Equal(t, setSeed(8), guessNumber(11))
}

func TestFirstBadVersion(t *testing.T) {
	setBadVersion(4)
	assert.Equal(t, 4, firstBadVersion(5))
	setBadVersion(1)
	assert.Equal(t, 1, firstBadVersion(5))
	assert.Equal(t, 1, firstBadVersion(1))
	setBadVersion(10)
	assert.Equal(t, 10, firstBadVersion(100))
}
