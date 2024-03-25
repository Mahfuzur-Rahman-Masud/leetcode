package s1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	obj := Constructor()
	obj.Push(2)
	assert.Equal(t, obj.Pop(), 2)
	assert.Equal(t, obj.Top(), 0)
	assert.True(t, obj.Empty())

	for i := 0; i < 560; i++ {
		obj.Push(i)
	}

	for i := 0; i < 560; i++ {
		assert.Equal(t, obj.Pop(), 559-i)
	}

}
