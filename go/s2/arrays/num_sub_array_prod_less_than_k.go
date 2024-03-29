package arrays

func numSubarrayProductLessThanK(nums []int, k int) int {

	c := 0
	for i := 0; i < len(nums); i++ {
		cum := 1
		for j := i; j < len(nums); j++ {

			cum *= nums[j]

			if cum < k {
				c++
			} else {
				break
			}
		}
	}

	return c
}
