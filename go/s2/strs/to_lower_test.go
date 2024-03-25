package strs

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestToLower(t *testing.T) {
	assert.Equal(t, "a", toLowerCase("A"))
	assert.Equal(t, "z", toLowerCase("Z"))
	assert.Equal(t, "c", toLowerCase("C"))
	assert.Equal(t, "hello", toLowerCase("HellO"))

}
