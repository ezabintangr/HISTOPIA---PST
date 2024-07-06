package main

import "fmt"

func highestPalindrome(s string, l int, r int, k int) string {
	if l >= r {
		return s
	}

	if s[l] != s[r] {
		maxChar := max(s[l], s[r])
		if maxChar > s[l] {
			s = replaceAtIndex(s, l, maxChar)
		} else {
			s = replaceAtIndex(s, r, maxChar)
		}
		k--
	}

	return highestPalindrome(s, l+1, r-1, k)
}

func max(a, b byte) byte {
	if a > b {
		return a
	}
	return b
}

func replaceAtIndex(s string, index int, newChar byte) string {
	bytes := []byte(s)
	bytes[index] = newChar
	return string(bytes)
}

func main() {
	s := "932139"
	k := 2

	result := highestPalindrome(s, 0, len(s)-1, k)
	fmt.Println(result)
}
