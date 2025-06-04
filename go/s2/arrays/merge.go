package arrays

func merge(nums1 []int, m int, nums2 []int, n int) {
	r := make([]int, m)
	ln := m + n
	i2, ir, li2, lir := 0, 0, n-1, m-1

	for i := 0; i < ln; i++ {
		if i < m {
			r[i] = nums1[i]
		}

		if ir > lir {
			nums1[i] = nums2[i2]
			i2++
		} else if i2 > li2 {
			nums1[i] = r[ir]
			ir++
		} else if nums2[i2] <= r[ir] {
			nums1[i] = nums2[i2]
			i2++
		} else {
			nums1[i] = r[ir]
			ir++
		}
	}

}
