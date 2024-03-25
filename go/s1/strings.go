package s1

import "strings"

func reverseWords(s string) string {
	b := []byte(s)
	markIn, markOut := 0, 1
	len := len(b)
	end := len - 1

	for markIn < markOut && markOut < len {
		if b[markOut] == ' ' || markOut == end {
			mi, mo := markIn, markOut-1
			if markOut == end {
				mo++
			}
			for mi < mo {
				b[mi], b[mo] = b[mo], b[mi]

				mi++
				mo--
			}
			markOut++
			markIn = markOut
		}

		markOut++
	}

	return string(b)
}

/**
You are given a license key represented as a string s that consists of only alphanumeric characters and dashes. The string is separated into n + 1 groups by n dashes. You are also given an integer k.

We want to reformat the string s such that each group contains exactly k characters, except for the first group, which could be shorter than k but still must contain at least one character. Furthermore, there must be a dash inserted between two groups, and you should convert all lowercase letters to uppercase.

Return the reformatted license key.



Example 1:

Input: s = "5F3Z-2e-9-w", k = 4
Output: "5F3Z-2E9W"
Explanation: The string s has been split into two parts, each part has 4 characters.
Note that the two extra dashes are not needed and can be removed.
Example 2:

Input: s = "2-5g-3-J", k = 2
Output: "2-5G-3J"
Explanation: The string s has been split into three parts,
each part has 2 characters except the first part as it could be shorter as mentioned above.
**/

func licenseKeyFormatting(s string, k int) string {
	b := []byte(s)
	l := len(b)

	xl := l + l/k
	o := make([]byte, xl)
	x := xl - 1
	xc := 0

	for i := l - 1; i >= 0; i-- {
		c := &b[i]
		if *c > 96 && *c < 123 {
			*c = *c - 32
		}

		if *c == '-' {
			continue
		}

		if xc == k {
			o[x] = '-'
			xc = 0
			x--
		}

		o[x] = *c
		x--
		xc++

	}

	return string(o[x+1:])

}

// Given a string s, check if it can be constructed by taking a substring of it and appending multiple copies of the substring together.
func repeatedSubstringPattern(s string) bool {
	return strings.Contains(((s + s)[1 : len(s)*2-1]), s)
	// return strings.Contains(ss, s)
}

// Given a string s which consists of lowercase or uppercase letters, return the length of the longest palindrome that can be built with those letters.
// "abccccdd" -> One longest palindrome that can be built is "dccaccd", whose length is 7.
func longestPalindrome(s string) int {
	max := 0
	sum := 0
	counter := [123]int{}

	for _, c := range s {
		counter[c]++
	}

	for i := 65; i < 123; i++ {
		b := counter[i]
		if b == 0 {
			continue
		}

		if b > max {
			if max%2 != 0 {
				sum--
			}

			max = b
			sum += max

		} else if b%2 == 0 {
			sum += b
		} else {
			sum += b - 1
		}
	}

	if len(s) > sum && max%2 == 0 {
		sum++
	}

	// z := counter[:]
	// sort.Ints(z)
	// fmt.Println(z)

	return sum
}

func reverseVowels(s string) string {
	i, i2 := 0, len(s)-1

	b := []byte(s)

	for i < i2 {

		switch b[i] {
		case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
			switch b[i2] {
			case 'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U':
				b[i], b[i2] = b[i2], b[i]
				i++
				i2--
			default:
				i2--
			}
		default:
			i++
		}

	}

	return string(b)
}

func isAnagram(s string, t string) bool {
	l := len(s)
	if l != len(t) {
		return false
	}

	c := [26]int16{}

	for i := 0; i < l; i++ {
		c[s[i]-'a']++
		c[t[i]-'a']--
	}

	for i := 0; i < 26; i++ {
		if c[i] != 0 {
			return false
		}
	}

	return true
}

func IsIsomorphic(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}

	st := [200]int{}
	ts := [200]int{}

	bs := []byte(s)
	bt := []byte(t)

	for i := 0; i < len(s); i++ {

		s := bs[i]
		t := bt[i]

		if st[s] != ts[t] {
			return false
		}

		st[s] = i + 1
		ts[t] = i + 1
	}

	return true
}

// func wordPattern(pattern string, s string) bool {
// 	words := strings.Split(s, " ")
// 	l := len(pattern)

// 	if len(words) != l {
// 		return false
// 	}

// 	wToChar := map[string]rune{}
// 	buff := [27]int{}

// 	for i, c := range pattern {
// 		if buff[c-96] == 0 {

// 		}
// 	}
// }
