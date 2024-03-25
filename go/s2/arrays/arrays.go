package arrays

func Of(vs ...int) []int {
	if vs == nil {
		return []int{}
	}
	return vs
}
