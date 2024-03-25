package s1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkConvertToBase72(t *testing.B) {

	for i := 0; i < 500; i++ {

		assert.Equal(t, "10", convertToBase72(7))
		assert.Equal(t, "11", convertToBase72(8))
		assert.Equal(t, "-11", convertToBase72(-8))
		assert.Equal(t, "-10", convertToBase72(-7))
		assert.Equal(t, "0", convertToBase72(0))
		assert.Equal(t, "1", convertToBase72(1))
		assert.Equal(t, "6", convertToBase72(6))
		assert.Equal(t, "202", convertToBase72(100))
	}
}

func BenchmarkConvertToBase7(t *testing.B) {

	for i := 0; i < 500; i++ {

		assert.Equal(t, "10", convertToBase7(7))
		assert.Equal(t, "11", convertToBase7(8))
		assert.Equal(t, "-11", convertToBase7(-8))
		assert.Equal(t, "-10", convertToBase7(-7))
		assert.Equal(t, "0", convertToBase7(0))
		assert.Equal(t, "1", convertToBase7(1))
		assert.Equal(t, "6", convertToBase7(6))
		assert.Equal(t, "202", convertToBase7(100))
	}
}

func BenchmarkHashInt(t *testing.B) {
	fmt.Println()
}

func BenchmarkReverseDigits(t *testing.B) {
	assert.Equal(t, 321, reverse(123))
	assert.Equal(t, 5, reverse(500))
	assert.Equal(t, 1, reverse(1))
	assert.Equal(t, 9, reverse(9))
	assert.Equal(t, 0, reverse(2147483647))
	assert.Equal(t, 0, reverse(-2147483648))
	assert.Equal(t, -412, reverse(-214))
}

func BenchmarkFindComplement(t *testing.B) {
	assert.Equal(t, 2, findComplement(5))
	assert.Equal(t, 0, findComplement(1))
	assert.Equal(t, 0, findComplement(0x7FFFFFFF))
}

func BenchmarkFlipBits(t *testing.B) {
	// num := uint8(42)
	// num2 := ^num
	// fmt.Printf("Binary representation of %d: %08b => %08b [%d]\n", num, num, num2, int8(num2))
	// var n int8 = 0b1101011
	// fmt.Println(n)

	// findComplement(2)
	fmt.Println(0x7FFFFFFF)
	fmt.Println(2147483647 == 0x7FFFFFFF)
	// fmt.Println(int32(1) << 31)
}

func TestToHex(t *testing.T) {
	fmt.Println(toHex(26))

	assert.Equal(t, "f", toHex(15))
	assert.Equal(t, "1", toHex(1))
	assert.Equal(t, "a", toHex(10))
	assert.Equal(t, "b", toHex(11))
	assert.Equal(t, "10", toHex(16))
	assert.Equal(t, "11", toHex(17))
	assert.Equal(t, "ffffffff", toHex(-1))
	assert.Equal(t, "0", toHex(0))
}

func TestIsPerfectSquare(t *testing.T) {
	assert.True(t, isPerfectSquare(1))
	assert.True(t, isPerfectSquare(4))
	assert.True(t, isPerfectSquare(9))
	assert.True(t, isPerfectSquare(16))
	assert.True(t, isPerfectSquare(25))
	assert.True(t, isPerfectSquare(808201))
	assert.True(t, !isPerfectSquare(8))
}

func TestIsPowerOf3(t *testing.T) {
	assert.True(t, isPowerOfThree(3))
	assert.True(t, isPowerOfThree(9))
	assert.True(t, isPowerOfThree(27))
	assert.True(t, isPowerOfThree(81))
	assert.True(t, !isPowerOfThree(1))
	assert.True(t, !isPowerOfThree(0))
	assert.True(t, !isPowerOfThree(12))
}

func BenchmarkIsPowerOf3(t *testing.B) {
	isPowerOfThree(3486784401)
}

func BenchmarkIsPowerOf32(t *testing.B) {
	isPowerOfThree2(3486784401)
}

func TestMissingNumber(t *testing.T) {
	assert.Equal(t, missingNumber([]int{3, 0, 1}), 2)
	assert.Equal(t, missingNumber([]int{0, 1}), 2)
	assert.Equal(t, missingNumber([]int{9, 6, 4, 2, 3, 5, 7, 0, 1}), 8)
}

func TestIsUgly(t *testing.T) {
	assert.True(t, isUgly_ugly_solution(6))
	assert.True(t, !isUgly_ugly_solution(7))
	assert.True(t, isUgly_ugly_solution(1))
	assert.True(t, !isUgly_ugly_solution(14))
	assert.True(t, isUgly_ugly_solution(25))
	assert.True(t, !isUgly_ugly_solution(13))
	assert.True(t, !isUgly_ugly_solution(39))
	assert.True(t, !isUgly_ugly_solution(-2147483648))
	assert.True(t, !isUgly_ugly_solution(937351770))
}

func TestIsHappy(t *testing.T) {
	m := map[int]bool{
		19: true,
		2:  false,
	}

	test(t, m, isHappy)
}

func TestAddDigits(t *testing.T) {
	assert.Equal(t, addDigits(38), 2)
	assert.Equal(t, addDigits(11), 2)
	assert.Equal(t, addDigits(0), 0)
	assert.Equal(t, addDigits(555), 6)
	assert.Equal(t, addDigits(55), 1)
}
