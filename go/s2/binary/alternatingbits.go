package binary

var x int32 = 0x55555555
var y int32 = 0x2aaaaaaa

func hasAlternatingBits(n int) bool {

	set := n&1 == 1
	for n > 0 {
		n >>= 1
		if (n&1 == 1) == set {
			return false
		}
		set = !set
	}

	return true
}

func hasAlternatingBits0_fail(n int) bool {
	return n|0x55555555 == 0x55555555 || n|0x2aaaaaaa == 0x2aaaaaaa
}
