package s1

func Merge(nums1 []int, m int, nums2 []int, n int) {
	buf := make([]int, m)

	rb, r2 := 0, 0

	for i := 0; i < m+n; i++ {
		if i < m {
			buf[i] = nums1[i]
		}

		if rb > m-1 {
			nums1[i] = nums2[r2]
			r2++
		} else if r2 > n-1 {

			nums1[i] = buf[rb]
			rb++

		} else if buf[rb] <= nums2[r2] {
			nums1[i] = buf[rb]
			rb++
		} else {
			nums1[i] = nums2[r2]
			r2++
		}
	}
}
