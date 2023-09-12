package leet

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func test[a comparable, b comparable](t *testing.T, m map[a]b, fn func(a) b) {
	fmt.Printf("%-24s\t\t%-24s\t\t%-24s\t\t%-8s\n", "INPUT", "EXPECT", "ACTUAL", "OK")
	fmt.Println("----------------------------------------------------------------------------------------------------------------------------")

	pass := true
	for i, o := range m {
		r := fn(i)
		fmt.Printf("%-24v\t\t%-24v\t\t%-24v\t\t%-8v\n", i, o, r, r == o)
		pass = pass && o == r

	}

	assert.Equal(t, pass, true)
}

type Case2[a comparable, b comparable, c comparable] struct {
	Input1   a
	Input2   b
	Expected c
	Actual   c
}

func newCase[a, b, c comparable](input1 a, input2 b, expected c) Case2[a, b, c] {
	return Case2[a, b, c]{
		Input1:   input1,
		Input2:   input2,
		Expected: expected,
	}
}

func testCase2[a, b, c comparable](t *testing.T, cs []*Case2[a, b, c], fn func(a, b) c) {
	pass := true

	fmt.Printf("%-24s\t\t%-24s\t\t%-24s\t\t%-8s\n", "INPUT", "EXPECT", "ACTUAL", "OK")
	fmt.Println("----------------------------------------------------------------------------------------------------------------------------")
	for _, c := range cs {
		c.Actual = fn(c.Input1, c.Input2)
		fmt.Printf("%-11v | %-11v\t\t%-24v\t\t%-24v\t\t%-8v\n", c.Input1, c.Input2, c.Expected, c.Actual, c.Actual == c.Expected)
		pass = pass && c.Expected == c.Actual
	}

	fmt.Println()
	assert.Equal(t, pass, true)
}
