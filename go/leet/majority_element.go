package leet

// majority vote is at least > len(nums) /2
func majorityMooreVotingAlg(nums []int) int {
	vote := 0
	candidate := 0

	for _, n := range nums {
		if vote == 0 {
			candidate = n
		}

		if candidate == n {
			vote++
		} else {
			vote--
		}
	}

	return candidate
}

func majorityElement(nums []int) int {
	max := 0
	m := map[int]int{}

	for _, n := range nums {
		m[n] = m[n] + 1
	}

	r := 0
	for k, v := range m {
		if v > max {
			max = v
			r = k
		}
	}

	return r
}
