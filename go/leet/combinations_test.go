package leet

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParenthesis(t *testing.T) {
	fmt.Println(generateParenthesis(3))

	assert.Equal(t, []string{}, generateParenthesis(0))
	assert.Equal(t, []string{"()"}, generateParenthesis(1))
	assert.Equal(t, []string{"(())", "()()"}, generateParenthesis(2))

	assert.Equal(t, []string{"((()))", "(()())", "(())()", "()(())", "()()()"}, generateParenthesis(3))

	assert.Equal(t, []string{"(((())))", "((()()))", "((())())", "((()))()", "(()(()))", "(()()())", "(()())()", "(())(())", "(())()()", "()((()))", "()(()())", "()(())()", "()()(())", "()()()()"}, generateParenthesis(4))
}
