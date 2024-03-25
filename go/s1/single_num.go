package s1

func singleNumber(nums []int) int {
	m := map[int]bool{}
	n := 0

	for i := 0; i < len(nums); i++ {
		n = nums[i]
		_, ok := m[n]
		if ok {
			delete(m, n)
		} else {
			m[n] = true
		}

	}

	for k := range m {
		return k
	}
	return -1
}

func singleNumberXor(nums []int) int {
	mask := 0

	for _, n := range nums {
		mask ^= n
	}

	return mask
}
