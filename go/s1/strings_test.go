package s1

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReverseWords(t *testing.T) {
	assert.Equal(t, "s'teL ekat edoCteeL tsetnoc", reverseWords("Let's take LeetCode contest"))
	assert.Equal(t, "doG gniD", reverseWords("God Ding"))
	assert.Equal(t, "gniD", reverseWords("Ding"))
	assert.Equal(t, "D", reverseWords("D"))
	assert.Equal(t, "CD", reverseWords("DC"))
	assert.Equal(t, "uoY ekil em", reverseWords("You like me"))
}

func TestLicensePlateFormat(t *testing.T) {
	fmt.Println(licenseKeyFormatting("abcdEERRxed", 3))
	assert.Equal(t, "A-BC-DE-ER-RX-ED", licenseKeyFormatting("abcdEERRxed", 2))
	assert.Equal(t, "AB-CDE-ERR-XED", licenseKeyFormatting("abcdEERRxed", 3))
	assert.Equal(t, "ABC-DEER-RXED", licenseKeyFormatting("abcdEERRxed", 4))
	assert.Equal(t, "A", licenseKeyFormatting("a", 3))
}

func TestRepeatedSubstringPattern(t *testing.T) {
	assert.True(t, repeatedSubstringPattern("abab"))
	assert.True(t, !repeatedSubstringPattern("aba"))
	assert.True(t, !repeatedSubstringPattern("abac"))
	assert.True(t, repeatedSubstringPattern("aaaaa"))
	assert.True(t, repeatedSubstringPattern("aa"))
	assert.True(t, !repeatedSubstringPattern("a"))
	assert.True(t, repeatedSubstringPattern("abcabcabcabc"))
	assert.True(t, repeatedSubstringPattern("dogdobdogdob"))
	assert.True(t, repeatedSubstringPattern("dogdogdog"))
	assert.True(t, repeatedSubstringPattern("milkmilkmilkmilk"))
	assert.True(t, repeatedSubstringPattern("honeyhoney"))
}

func TestLongestPalindrome(t *testing.T) {
	assert.Equal(t, 1, longestPalindrome("a"))
	assert.Equal(t, 7, longestPalindrome("abccccdd"))
	assert.Equal(t, 1, longestPalindrome("a"))
	assert.Equal(t, 3, longestPalindrome("aac"))
	assert.Equal(t, 3, longestPalindrome("aaac"))
	assert.Equal(t, 5, longestPalindrome("aaacc"))
	assert.Equal(t, 4, longestPalindrome("aacc"))
	assert.Equal(t, 5, longestPalindrome("aaccx"))
	assert.Equal(t, 55, longestPalindrome("zeusnilemacaronimaisanitratetartinasiaminoracamelinsuez"))
	assert.Equal(t, 983, longestPalindrome("civilwartestingwhetherthatnaptionoranynartionsoconceivedandsodedicatedcanlongendureWeareqmetonagreatbattlefiemldoftzhatwarWehavecometodedicpateaportionofthatfieldasafinalrestingplaceforthosewhoheregavetheirlivesthatthatnationmightliveItisaltogetherfangandproperthatweshoulddothisButinalargersensewecannotdedicatewecannotconsecratewecannothallowthisgroundThebravelmenlivinganddeadwhostruggledherehaveconsecrateditfaraboveourpoorponwertoaddordetractTgheworldadswfilllittlenotlenorlongrememberwhatwesayherebutitcanneverforgetwhattheydidhereItisforusthelivingrathertobededicatedheretotheulnfinishedworkwhichtheywhofoughtherehavethusfarsonoblyadvancedItisratherforustobeherededicatedtothegreattdafskremainingbeforeusthatfromthesehonoreddeadwetakeincreaseddevotiontothatcauseforwhichtheygavethelastpfullmeasureofdevotionthatweherehighlyresolvethatthesedeadshallnothavediedinvainthatthisnationunsderGodshallhaveanewbirthoffreedomandthatgovernmentofthepeoplebythepeopleforthepeopleshallnotperishfromtheearth"))
}

func TestReverseVowels(t *testing.T) {
	assert.Equal(t, "holle", reverseVowels("hello"))
	assert.Equal(t, "leotcede", reverseVowels("leetcode"))
}

func TestIsAnagram(t *testing.T) {
	assert.True(t, isAnagram("anagram", "nagaram"))
	assert.True(t, isAnagram("anag", "gana"))
	assert.True(t, !isAnagram("ana", "ann"))
	assert.True(t, !isAnagram("ana", "gana"))
	assert.True(t, !isAnagram("cat", "rat"))
}

func TestIsIsomorphic(t *testing.T) {
	cs := []*Case2[string, string, bool]{
		{
			Input1:   "egg",
			Input2:   "add",
			Expected: true,
		},
		{
			Input1:   "foo",
			Input2:   "bar",
			Expected: false,
		},

		{
			Input1:   "paper",
			Input2:   "title",
			Expected: true,
		},

		{
			Input1:   "badc",
			Input2:   "baba",
			Expected: false,
		},
	}
	testCase2[string, string, bool](t, cs, IsIsomorphic)
}
