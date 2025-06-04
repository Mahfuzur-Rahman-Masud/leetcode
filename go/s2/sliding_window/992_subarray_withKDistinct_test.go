package slidingwindow

import (
	"leets/s2/arrays"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSubarraysWithKDistinct(t *testing.T) {
	// assert.Equal(t, 7, subarraysWithKDistinct(arrays.Of(1, 2, 1, 2, 3), 2))
	// assert.Equal(t, 3, subarraysWithKDistinct(arrays.Of(1, 2, 1, 3, 4), 3))
	// assert.Equal(t, 8, subarraysWithKDistinct(arrays.Of(2, 1, 1, 1, 2), 1)) // 2, 1, 1, 1, 2 , 11, 111, 11

	assert.Equal(t, 149, subarraysWithKDistinct(arrays.Of(27, 27, 43, 28, 11, 20, 1, 4, 49, 18, 37, 31, 31, 7, 3, 31, 50, 6, 50, 46, 4, 13, 31, 49, 15, 52, 25, 31, 35, 4, 11, 50, 40, 1, 49, 14, 46, 16, 11, 16, 39, 26, 13, 4, 37, 39, 46, 27, 49, 39, 49, 50, 37, 9, 30, 45, 51, 47, 18, 49, 24, 24, 46, 47, 18, 46, 52, 47, 50, 4, 39, 22, 50, 40, 3, 52, 24, 50, 38, 30, 14, 12, 1, 5, 52, 44, 3, 49, 45, 37, 40, 35, 50, 50, 23, 32, 1, 2), 20))
}

func subarraysWithKDistinct(nums []int, k int) int {
	return subarraysWithAtMostKDistinct(nums, k) - subarraysWithAtMostKDistinct(nums, k-1)
}

func subarraysWithAtMostKDistinct(nums []int, k int) int {
	if k == 0 || len(nums) < k {
		return 0
	}

	out := 0
	// start := 0
	left := 0
	count := 0
	countOccur := map[int]int{}

	for right, n := range nums {
		countOccur[n]++

		if countOccur[n] == 1 {
			count++
		}

		for count > k {
			n2 := nums[left]
			countOccur[n2]--
			if countOccur[n2] == 0 {
				count--
			}

			left++
		}
		out += right - left + 1

		// if l == k {
		// 	for i2 := mark; i2 <= i; i2++ {
		// 		n2 := nums[i2]
		// 		m[n2]--

		// 		if m[n2] <= 0 {
		// 			mark = i2
		// 			m[n2] = 1
		// 			break
		// 		}
		// 	}

		// 	out += mark - start + 1
		// } else if l > k {

		// 	n2 := nums[mark]
		// 	delete(m, n2)
		// 	mark++
		// 	start = mark
		// 	out++
		// }

	}

	return out
}
