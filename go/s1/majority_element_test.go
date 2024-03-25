package s1

import (
	"fmt"
	"testing"
)

func TestMajorityElement(t *testing.T) {
	n := []int{3, 2, 3}
	fmt.Println(majorityElement(n))
	fmt.Println(majorityElement([]int{2, 2, 1, 1, 1, 2, 2}))
	fmt.Println(majorityMooreVotingAlg([]int{2, 2, 1, 1, 1, 2, 2}))
}
