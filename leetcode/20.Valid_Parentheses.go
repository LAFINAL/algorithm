package leetcode

func isValid(s string) bool {

	match := make(map[byte]byte)
	match[']'] = '['
	match['}'] = '{'
	match[')'] = '('

	stack := make([]byte, 0, 0)
	for i:=0; i<len(s); i++ {
		if match[s[i]] == 0 { // map 안에 없으면.
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}

			pop := stack[len(stack)-1]
			if pop != match[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}
