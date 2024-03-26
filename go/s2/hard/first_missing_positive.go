package hard

func firstMissingPositive(nums []int) int {
	p := [100002]int{} // constant space
	l := len(nums)

	for _, n := range nums {
		if n < 1 || n > l {
			continue
		}

		p[n] = 1
	}

	for i := 1; i < 100001; i++ {
		if p[i] == 0 {
			return i
		}
	}

	return l + 1
}

// constant space
func firstMissingPositive1(nums []int) int {
	l := len(nums)
	for i := 1; i < l; i++ {

		for nums[i] > 0 && nums[i] <= l && nums[nums[i]-1] != nums[i] {
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}

	for i, n := range nums {
		if n != i+1 {
			return i + 1
		}
	}

	return l + 1
}

func firstMissingPositive0(nums []int) int {
	if nums == nil {
		return 1
	}

	l := len(nums)
	pn := make([]int, l+2)

	for _, n := range nums {
		if n < 1 || n > l+1 {
			continue
		}
		pn[n] = 1
	}

	for i, p := range pn {
		if p == 0 && i > 0 {
			return i
		}
	}

	return 1
}
