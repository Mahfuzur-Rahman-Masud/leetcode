package s1

import (
	"math"
	"strconv"
)

func convertToBase72(num int) string {
	if num == 0 {
		return "0"
	}

	negative := num < 0

	if num < 0 {
		num *= -1
	}

	b := make([]byte, 0, 64)

	for num > 0 {
		b = append(b, byte(num%7+48))
		num /= 7
	}

	if negative {
		b = append(b, '-')
	}

	l, r := 0, len(b)-1
	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}

	return string(b)
}

func convertToBase7(num int) string {
	if num == 0 {
		return "0"
	}

	sign := ""
	if num < 0 {
		sign = "-"
		num *= -1
	}

	// b := []byte{}
	out := ""

	for num > 0 {
		// b = append(b, byte(num%7))
		out = strconv.Itoa(num%7) + out
		num /= 7
	}

	// l, r := 0, len(b)-1
	// for l < r {
	// 	b[l], b[r] = b[r], b[l]
	// 	l++
	// 	r--
	// }

	// return string(b)
	return sign + out

}

//Given a signed 32-bit integer x, return x with its digits reversed. If reversing x causes the value to go outside the signed 32-bit integer range [-231, 231 - 1], then return 0.

// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).
func reverse(x int) int {
	b := []byte(strconv.Itoa(x))
	l, r := 0, len(b)-1

	if b[0] == '-' {
		l++
	}

	for l < r {
		b[l], b[r] = b[r], b[l]
		l++
		r--
	}

	ans, e := strconv.ParseInt(string(b), 10, 32)
	if e != nil {
		return 0
	}
	return int(ans)
}

// convert zeroes to one and ones to zeroes
// eliminate all leading zeroes first
func findComplement(num int) int {
	u := uint32(num)
	b := 0
	umod, uprev := u, u
	umod <<= 1

	for umod>>1 == uprev {
		umod, uprev = umod<<1, umod
		b++
	}

	return int((^u) << b >> b)
}

func toHex(num int) string {
	if num == 0 {
		return "0"
	}

	code := []byte{'0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'}
	out := make([]byte, 0, 64)
	nn := uint32(num)

	for nn > 0 {
		n := nn % 16
		out = append(out, code[n])
		nn /= 16
	}

	for i, j := 0, len(out)-1; i < j; i, j = i+1, j-1 {
		out[i], out[j] = out[j], out[i]
	}

	return string(out)
}

func isPerfectSquare(num int) bool {
	l, h := 0, num/2

	for l <= h {
		mid := l + (h-l)/2
		r := mid * mid
		if r == num {
			return true
		} else if r > num {
			h = mid - 1
		} else {
			l = mid + 1
		}
	}

	return l*l == num
}

func missingNumber(nums []int) int {
	sum := 0

	for _, n := range nums {
		sum += n
	}

	l := len(nums)
	return (l+1)*(l)/2 - sum
}

func isUgly(number int) bool {
	if number < 1 {
		return false
	}

	if number == 1 {
		return true
	}

	c := make(chan int, 3)

	for _, factor := range [3]int{2, 3, 5} {

		go func(f int, n int) {
			for ; n%f == 0; n /= f {
			}
			c <- n
		}(factor, number)
	}

	return <-c*<-c*<-c == number*number
}

// times out
func isUgly_ugly_solution(n int) bool {
	if n < 1 {
		return false
	}
	if n < 0 {
		n = n * -1
	}

	if n < 7 {
		return true
	}

	if n%7 == 0 {
		return false
	}

	x := 11

outer:
	for x <= n {

		// is x prime?
		if x%2 == 0 || x%3 == 0 {
			x++
			continue
		}

		i := 5
		sqrt := int(math.Sqrt(float64(x)))

		for i <= sqrt {
			if x%i == 0 || x%(i+2) == 0 {
				x++
				continue outer
			}

			i += 5
		}

		if n%x == 0 {
			return false
		}
		x++
	}

	return true
}

func addDigits(num int) int {
	var next int

	for num > 9 {
		next = num
		num = 0
		for next > 0 {
			num += next % 10
			next /= 10
		}
	}

	return num
}
func isPowerOfFour(n int) bool {
	if n < 1 {
		return false
	}

	for n%4 == 0 {
		n /= 4
	}

	return n == 1
}

func isPowerOfThree3(n int) bool {
	if n < 1 {
		return false
	}

	for n%3 == 0 {
		n /= 3
	}

	return n == 1
}

func isPowerOfThree2(n int) bool {
	for n > 2 {
		n -= 3
	}

	return n == 0
}

func isPowerOfThree(n int) bool {
	if n < 1 {
		return false
	}

	if n&1 != 1 {
		return false
	}

	for n > 1 {
		if n%3 != 0 {
			return false
		}
		n /= 3
	}

	return true
}

func isHappy(n int) bool {
	slow, fast := n, sumSqrOfDigits(n)

	for fast != 1 && slow != fast {
		slow = sumSqrOfDigits(slow)
		fast = sumSqrOfDigits(sumSqrOfDigits(fast))
	}

	return fast == 1
}

func sumSqrOfDigits(n int) int {
	sum := 0
	for n != 0 {
		d := n % 10
		n /= 10
		sum += d * d
	}

	return sum
}
