package s1

// The Hamming distance between two integers is the number of positions at which the corresponding bits are different.
/**
Example 1:

Input: x = 1, y = 4
Output: 2
Explanation:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑
The above arrows point to positions where the corresponding bits are different.
Example 2:

Input: x = 3, y = 1
Output: 1
**/
func hammingDistance(x int, y int) int {
	xor := x ^ y
	count := 0

	// for xor != 0 {
	// 	count += xor & 1
	// 	xor >>= 1
	// }

	for xor > 0 {
		xor &= (xor - 1)
		count++
	}

	return count
}

func divideBy2(n int) int {
	return n >> 1
}

func countBits2(n int) []int {
	l := n + 1
	out := make([]int, l)

	for i := 1; i < l; i++ {
		if i&1 == 1 {
			out[i] = out[i-1] + 1
		} else {
			out[i] = out[i>>1]
		}
	}

	return out
}

func countBits(n int) []int {
	out := make([]int, n+1)

	for i := 0; i < n+1; i++ {
		count := 0
		x := i
		for x > 0 {
			x &= x - 1
			count++
		}

		out[i] = count
	}
	return out
}

func countSetBits(n int) int {
	count := 0
	for n > 0 {
		n = n & (n - 1)
		count++
	}
	return count
}

func reverseBits(n uint32) uint32 {
	n = (n&0xffff0000)>>16 | (n&0x0000ffff)<<16
	n = (n&0xff00ff00)>>8 | (n&0x00ff00ff)<<8
	n = (n&0xf0f0f0f0)>>4 | (n&0x0f0f0f0f)<<4
	n = (n&0xcccccccc)>>2 | (n&0x33333333)<<2
	n = (n&0xaaaaaaaa)>>1 | (n&0x55555555)<<1

	return n
}

func hammingWeight(n uint32) int {
	var w uint32 = 0

	for n > 0 {
		w += n & 1
		n >>= 1
	}

	return int(w)
}
