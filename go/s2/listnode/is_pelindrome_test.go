package listnode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPenlidrome(t *testing.T) {
	assert.True(t, isPalindrome(ListOf(1, 2, 2, 1)))
	assert.True(t, !isPalindrome(ListOf(1, 2)))
	assert.True(t, isPalindrome(ListOf(1)))
}
