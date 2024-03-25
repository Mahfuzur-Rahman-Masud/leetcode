package algo

func search(nums []int, target int) int {
	if nums == nil || len(nums) == 0 {
		return -1
	}

	l := len(nums)
	min, max := 0, l-1

	for min <= max {
		i := min + (max-min)/2
		cur := nums[i]
		if cur == target {
			return i
		}

		if cur < target {
			min = i + 1
		} else {
			max = i - 1
		}
	}

	return -1
}
