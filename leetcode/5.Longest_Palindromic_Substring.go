package leetcode

func LongestPalindrome(s string) string {

	for i:=len(s); i>0; i-- {
		for j:=0; j<len(s)-i+1; j++ {
			// j부터 +len(s)-1까지지
			if IsPalindrome(s[j:len(s)]) {
				return s[j:len(s)]
			}
		}
	}

	return ""
}

func IsPalindrome(s string) bool {
	r := make([]rune, len(s))
	for i, v := range s {
		r[len(s)-i-1] = v
	}
	if string(r) == s {
		return true
	}
	return false
}

// refactor.
//for i:=0; i<len(s)/2; i++ {
//if s[i] != s[len(s)-i-1] {
//return false
//}
//}
//return true