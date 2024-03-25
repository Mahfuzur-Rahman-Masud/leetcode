package s1

import (
	"fmt"
	"testing"
)

func TestClimbStairs(t *testing.T) {

	r := ClimbStairs(2)
	fmt.Println(r)
	fmt.Println(ClimbStairs(10), Fibonacchi(10), Fib(10))
}

func TestFiboncchi(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("F(%d): %d | %d | %d \n", i, Fib(i), ClimbStairs(i), Fib2(i))
	}
}

func TestFib2(t *testing.T) {
	for i := 0; i < 10; i++ {
		fmt.Printf("F(%d) : %d\n", i, Fib2(i))
	}
}
