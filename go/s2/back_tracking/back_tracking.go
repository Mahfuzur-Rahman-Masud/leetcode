package backtracking

func backTrack(nums []int, n int) int {
	res := n
	for _, v := range nums {
		print(v)
		print("\t")
	}
	println()

	for i, v := range nums {
		res += backTrack(nums[i+1:], n^v)
	}
	return res
}

func subsetXORSum(nums []int) int {
	return backTrack(nums, 0)
}
