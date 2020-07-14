package leetcode

import "fmt"

func longestPalindromeWithDP(s string) {
	a := make([][]bool, len(s), len(s))
	for i:=0; i<len(a); i++ {
		a[i] = make([]bool, len(s), len(s))
	}

	// length = 1
	for i:=0; i<len(s); i++ {
		a[i][i] = true
	}

	// length = 2
	for i:=0; i<len(s)-1; i++ {
		if s[i] == s[i+1] {
			a[i][i+1] = true
		}
	}
}

func longestPalindrome(s string) string {

	for i:=len(s); i>0; i-- {
		for j:=0; j<len(s)-i+1; j++ {
			a := isPalindrome2(s[j:j+i])
			if a == true {
				return s[j:j+i]
			}
		}
	}

	return ""
}

func isPalindrome2(s string) bool {
	fmt.Println(s)
	for i:=0; i<len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
