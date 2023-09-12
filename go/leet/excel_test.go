package leet

import (
	"fmt"
	"testing"
)

func TestToTitle(t *testing.T) {
	fmt.Println(convertToTitle(26))
	fmt.Println(convertToTitle(28))
	fmt.Println(convertToTitle(1))
	fmt.Println(convertToTitle(701))
}

func TestToColNumber(t *testing.T) {
	fmt.Println(titleToNumber("A"))
	fmt.Println(titleToNumber("Z"))
	fmt.Println(titleToNumber("ZY"))
	fmt.Println(titleToNumber("AB"))
}

func TestPow(t *testing.T) {
	fmt.Println(pow(2, 2))
	fmt.Println(pow(2, 0))
	fmt.Println(pow(0, 10000))
	fmt.Println(pow(2, 1))
	fmt.Println(pow(2, 3))
}
