package arrays

func findShortestSubArray(nums []int) int {
	x := [50000]int{}
	min := [50000]int{}
	max := [50000]int{}
	d := 0
	v := 0

	for i, n := range nums {
		if x[n] == 0 {
			min[n] = i
		}

		x[n]++
		max[n] = i
		if x[n] > d {
			d = x[n]
			v = n
		} else if x[n] == d && (i-min[n] < max[v]-min[v]) {
			v = n
		}
	}

	return max[v] - min[v] + 1
}
