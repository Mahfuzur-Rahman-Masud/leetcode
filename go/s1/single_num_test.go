package s1

import (
	"fmt"
	"testing"
)

func TestSingleNum(t *testing.T) {
	r := singleNumber([]int{2, 2, 1})
	fmt.Println(r)
	fmt.Println(singleNumber([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumber([]int{1}))
	fmt.Println(singleNumber([]int{5, 4, 4}))
}

func TestSingleNumXor(t *testing.T) {

	// fmt.Println(0 ^ 2 ^ 2 ^ 1)

	fmt.Println(singleNumberXor([]int{2, 2, 1}))
	fmt.Println(singleNumberXor([]int{4, 1, 2, 1, 2}))
	fmt.Println(singleNumberXor([]int{1}))
	fmt.Println(singleNumberXor([]int{5, 4, 4}))
}
