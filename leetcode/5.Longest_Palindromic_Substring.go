package leetcode

func LongestPalindrome(s string) string {
	max := 0
	answer := ""
	for i:=0; i<len(s); i++ {
		left := i-1
		right := i+1
		if right < len(s) && s[right] == s[right-1] && left >= 0 && s[left] != s[left+1]{
			right = right + 1
		}

		for {
			prevLength := (right-1) - (left+1) + 1
			if left < 0 || right > len(s) -1 {
				if prevLength > max {
					answer = s[left+1:right]
					max = prevLength
				}
				break
			}

			if s[left] == s[right] {
				left = left - 1
				right = right + 1
			} else {
				if prevLength > max {
					answer = s[left+1 : right]
					max = prevLength
				}
				break
			}
		}

	}
	return answer
}

