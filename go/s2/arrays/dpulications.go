package arrays

func findDuplicate(nums []int) int {
	x := [500001]int{}

	for _, n := range nums {
		x[n]++
		if x[n] > 1 {
			return n
		}
	}

	return -1
}

// Given an integer array nums of length n where all the integers of nums are in the range [1, n] and each integer appears once or twice, return an array of all the integers that appears twice.
// You must write an algorithm that runs in O(n) time and uses only constant extra space.
func findDuplicates(nums []int) []int {
	out := make([]int, 0, 256)
	dup := [500001]int{}

	for _, n := range nums {
		if dup[n] == 1 {
			out = append(out, n)
		}

		dup[n]++
	}

	return out
}
