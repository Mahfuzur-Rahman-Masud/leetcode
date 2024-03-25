package s1

import (
	"fmt"
	"regexp"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsMatch(t *testing.T) {
	r, e := regexp.Compile("^aa.*bc*$")

	if e != nil {
		e = fmt.Errorf("Pattern compile error: %s", e)
		fmt.Println(e)
		return
	}

	fmt.Printf("Matches %v\n", r.Match([]byte("aaaaaabc")))

	// assert.True(t, isMatch("aa", "aa"))
	// assert.True(t, isMatch("aa", "a*"))
	// assert.True(t, !isMatch("ab", "a*"))
	// assert.True(t, !isMatch("aa", "a**"))
	// assert.True(t, isMatch("aa", ".*"))
	// assert.True(t, isMatch("aaaaa", ".*"))
	// assert.True(t, isMatch("a", ".*"))
	// assert.True(t, isMatch("a", ".*"))
	// assert.True(t, isMatch("aa", ".*a"))
	assert.True(t, !isMatch("aa", ".*c"))

}
