package slidingwindow2

import (
	"leets/s2/arrays"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCountSubarrays(t *testing.T) {
	assert.Equal(t, int64(2), countSubarrays(arrays.Of(1, 3, 5, 2, 7, 5), 1, 5))
	assert.Equal(t, int64(3), countSubarrays(arrays.Of(1, 3, 5, 2, 7, 5, 1), 1, 5))
	assert.Equal(t, int64(10), countSubarrays(arrays.Of(1, 1, 1, 1), 1, 1))
	assert.Equal(t, int64(1), countSubarrays(arrays.Of(1), 1, 1))
	assert.Equal(t, int64(0), countSubarrays(arrays.Of(), 1, 1))
}

func countSubarrays(nums []int, minK int, maxK int) int64 {
	left := 0

	minC, maxC := 0, 0
	out := int64(0)

	for i, n := range nums {
		if n < minK || n > maxK {
			out += countSubarrays(nums[i+1:], minK, maxK)
			break
		}

		if n == minK {
			minC++
		}
		if n == maxK {
			maxC++
		}

		if minC < 1 || maxC < 1 {
			continue
		}

		for left <= i {
			n := nums[left]

			if n == minK {
				minC--
			}
			if n == maxK {
				maxC--
			}

			if minC == 0 || maxC == 0 {
				break
			}

			left++
		}

		if minC == 0 {
			minC = 1
		}
		if maxC == 0 {
			maxC = 1
		}

		out += int64(left) + 1
	}

	return out
}
