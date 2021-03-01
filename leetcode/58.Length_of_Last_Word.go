package leetcode

import "strings"

func lengthOfLastWord(s string) int {
	return len(s) - strings.LastIndex(s, " ") -1
}