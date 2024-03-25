package arrays

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindDuplicate(t *testing.T) {
	assert.Equal(t, 2, findDuplicate(Of(1, 3, 4, 2, 2)))
	assert.Equal(t, 3, findDuplicate(Of(1, 3, 4, 2, 3)))
}

func TestFindDuplicates(t *testing.T) {
	assert.Equal(t, Of(1), findDuplicates(Of(1, 1, 2)))
	assert.Equal(t, Of(), findDuplicates(Of(1)))
	assert.Equal(t, Of(2, 3), findDuplicates(Of(4, 3, 2, 7, 8, 2, 3, 1)))
}
