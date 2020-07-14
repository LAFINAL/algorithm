package leetcode

import "strconv"

func isPalindrome(x int) bool {

	if x < 0 {
		return false
	}

	if x == 0 {
		return true
	}

	revertedNumber := 0
	for x > revertedNumber {
		revertedNumber = revertedNumber * 10 + x%10
		x = x/10
	}

	return x == revertedNumber || x == revertedNumber/10
}

func isPalindromeWithString(x int) bool {

	if x < 0 {
		return false
	}
	stringX := strconv.Itoa(x)
	for i:=0; i<len(stringX)/2; i++ {
		if stringX[i] != stringX[len(stringX)-1-i] {
			return false
		}
	}

	return true
}


//func isPalindrome(x int) bool {
//
//	if x == 0 {
//		return true
//	}
//
//	if x < 0 || x%10 == 0 {
//		return false
//	}
//
//	revertedNumber := 0
//	for x > revertedNumber {
//		revertedNumber = revertedNumber * 10 + x % 10
//		x = x / 10
//	}
//
//	return x == revertedNumber || x == revertedNumber/10
//}