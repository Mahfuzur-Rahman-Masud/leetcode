package leet

func TwoSum(nums []int, target int) []int {
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				return []int{i, j}
			}
		}
	}

	return nil
}

func TwoSumWithMap(nums []int, target int) []int {
	m := map[int]int{}

	for i := 0; i < len(nums); i++ {
		val, ok := m[target-nums[i]]
		if ok {
			return []int{i, val}
		}

		m[nums[i]] = i
	}

	return nil

}
