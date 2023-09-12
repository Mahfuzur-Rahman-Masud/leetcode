package leet

import "testing"

func BenchmarkTwoSum(b *testing.B) {
	sample := []int{2, 7, 11, 15}
	target := 9

	r := TwoSum(sample, target)
	if sample[r[0]]+sample[r[1]] != target {
		b.Error("did not match", sample, target, r)
	}

}

func BenchmarkTwoSum2(b *testing.B) {
	sample := []int{2, 7, 11, 15}
	target := 9

	r := TwoSum(sample, target)
	if sample[r[0]]+sample[r[1]] != target {
		b.Error("did not match", sample, target, r)
	}

}
