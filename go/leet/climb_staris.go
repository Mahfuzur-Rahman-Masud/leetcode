package leet

func ClimbStairs(n int) int {
	prev, curr := 1, 1
	// var next int

	for i := 1; i < n; i++ {
		curr, prev = prev+curr, curr
		// next = prev + curr
		// prev = curr
		// curr = next
	}

	return curr

}

func Fibonacchi(i int) int {
	if i < 2 {
		return i
	}

	n_minus_2, n_minus_1 := 0, 1

	ans := 0

	for n := 0; n < i; n++ {
		ans = n_minus_1 + n_minus_2
		n_minus_2 = n_minus_1
		n_minus_1 = ans
	}

	return ans

}

func Fib2(n int) int {

	i, prev, curr := -1, -1, 1

	for i < n {
		prev, curr = curr, prev+curr
		i++
	}

	return curr
}

func Fib(n int) int {
	if n < 2 {
		return n
	}

	l2, l1 := 0, 1
	ans := 0

	for i := 2; i <= n; i++ {
		ans = l1 + l2
		l2 = l1
		l1 = ans
	}

	return ans
}
