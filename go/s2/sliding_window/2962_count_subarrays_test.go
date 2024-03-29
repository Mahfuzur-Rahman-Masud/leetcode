package slidingwindow

import (
	"leets/s2/arrays"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSubarrays(t *testing.T) {
	assert.Equal(t, int64(6), countSubarrays(arrays.Of(1, 3, 2, 3, 3), 2))
	assert.Equal(t, int64(0), countSubarrays(arrays.Of(1, 4, 2, 1), 3))
	assert.Equal(t, int64(1), countSubarrays(arrays.Of(1), 1))
	assert.Equal(t, int64(2), countSubarrays(arrays.Of(1, 2), 1))
	// assert.Equal(t, int64(1), countSubarrays(arrays.Of(1, 4, 2, 1), 1))
}

func countSubarrays(nums []int, k int) int64 {
	mx := 0
	for _, n := range nums {
		if n > mx {
			mx = n
		}
	}

	out := int64(0)
	cur := 0
	ln := len(nums)
	trace := make([]int, 0, ln/2)

	for cur < ln {
		n := nums[cur]
		if n == mx {
			trace = append(trace, cur)
		}

		l := len(trace)
		if l >= k {
			out += int64(trace[l-k]) + 1
		}

		cur++
	}

	return out
}
