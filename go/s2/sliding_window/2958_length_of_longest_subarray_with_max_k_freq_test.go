package slidingwindow

import (
	"leets/s2/arrays"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubarrayLength(t *testing.T) {
	assert.Equal(t, 6, maxSubarrayLength(arrays.Of(1, 2, 3, 1, 2, 3, 1, 2), 2))
	assert.Equal(t, 2, maxSubarrayLength(arrays.Of(1, 2, 1, 2, 1, 2, 1, 2), 1))
	assert.Equal(t, 4, maxSubarrayLength(arrays.Of(5, 5, 5, 5, 5, 5, 5), 4))
	assert.Equal(t, 0, maxSubarrayLength(arrays.Of(5, 5, 5, 5, 5, 5, 5), 0))
	assert.Equal(t, 1, maxSubarrayLength(arrays.Of(5, 5, 5, 5, 5, 5, 5), 1))
	assert.Equal(t, 1, maxSubarrayLength(arrays.Of(1), 1))
}

func maxSubarrayLength(nums []int, k int) int {
	l := len(nums)
	if l == 0 || k == 0 {
		return 0
	}

	st, i := 0, 0
	tally := map[int]int{}
	v := 0
	ll := 0
	ml := 0

	for i < len(nums) {
		v = nums[i]

		if tally[v] != k {
			v = nums[i]
			ll = i - st + 1

			if ll > ml {
				ml = ll
			}
		} else {

			for j := st; j < i; j++ {
				vv := nums[j]
				tally[vv]--
				if tally[vv] == 0 {
					delete(tally, vv)
				}

				if vv == v {
					st = j + 1
					if l-st <= ml {
						return ml
					}
					break
				}

			}
		}

		tally[v]++
		i++
	}

	return ml
}
