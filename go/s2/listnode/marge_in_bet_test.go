package listnode

import (
	"fmt"
	"testing"
	"time"
)

func TestMergeInBetween(t *testing.T) {
	l1 := CreateLLFrom([]int{0, 1, 2, 3, 4, 5, 6})
	l2 := CreateLLFrom([]int{1000, 1001, 1002, 1003, 1004})

	out := mergeInBetween(l1, 2, 5, l2).ToString()
	fmt.Println(out)

	l1 = CreateLLFrom([]int{0, 1, 2, 3, 4, 5, 6})
	l2 = CreateLLFrom([]int{1000, 1001, 1002, 1003, 1004})

	fmt.Println(mergeInBetween(l1, 2, 2, l2).ToString())

	l1 = CreateLLFrom([]int{0, 1, 2, 3, 4, 5, 6})
	l2 = CreateLLFrom([]int{1000, 1001, 1002, 1003, 1004})

	fmt.Println(mergeInBetween(l1, 5, 5, l2).ToString())

	l1 = CreateLLFrom([]int{0, 1, 2, 3, 4, 5, 6})
	l2 = CreateLLFrom([]int{1000, 1001, 1002, 1003, 1004})

	fmt.Println(mergeInBetween(l1, 1, 1, l2).ToString())

	l1 = CreateLLFrom([]int{10, 1, 13, 6, 9, 5})
	l2 = CreateLLFrom([]int{1000000, 1000001, 1000002})

	fmt.Println(mergeInBetween(l1, 3, 4, l2).ToString())

}

func TestXxx(t *testing.T) {
	a := time.Now().UnixMilli()
	for i := 0; i < 1000; i++ {
		l1 := CreateLLFrom([]int{0, 1, 2, 3, 4, 5, 6})
		l2 := CreateLLFrom([]int{1000, 1001, 1002, 1003, 1004})
		fmt.Println(mergeInBetween(l1, 2, 5, l2).ToString())
		l1 = CreateLLFrom([]int{0, 1, 2, 3, 4, 5, 6})
		l2 = CreateLLFrom([]int{1000, 1001, 1002, 1003, 1004})
		fmt.Println(mergeInBetween(l1, 2, 2, l2).ToString())

		l1 = CreateLLFrom([]int{0, 1, 2, 3, 4, 5, 6})
		l2 = CreateLLFrom([]int{1000, 1001, 1002, 1003, 1004})
		fmt.Println(mergeInBetween(l1, 5, 5, l2).ToString())

		l1 = CreateLLFrom([]int{0, 1, 2, 3, 4, 5, 6})
		l2 = CreateLLFrom([]int{1000, 1001, 1002, 1003, 1004})
		fmt.Println(mergeInBetween(l1, 1, 1, l2).ToString())

		l1 = CreateLLFrom([]int{10, 1, 13, 6, 9, 5})
		l2 = CreateLLFrom([]int{1000000, 1000001, 1000002})
		fmt.Println(mergeInBetween(l1, 3, 4, l2).ToString())
	}

	fmt.Println(time.Now().UnixMilli() - a)

}
