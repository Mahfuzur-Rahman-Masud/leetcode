package strs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateString(t *testing.T) {
	assert.True(t, rotateString("abcde", "cdeab"))
	assert.False(t, rotateString("abcde", "abced"))
}
