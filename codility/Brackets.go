package codility

// Brackets는 stack과 Map을 활용한다.
func Brackets(S string) int {
	pare := make(map[byte]byte)
	pare[']'] = '['
	pare['}'] = '{'
	pare[')'] = '('

	stack := make([]byte, 0, 0)

	for i:=0; i<len(S); i++ {
		value, exist := pare[S[i]]
		if !exist {
			stack = append(stack, S[i])
			continue
		}

		if len(stack) == 0 {
			return 0
		}
		pop := stack[len(stack)-1]
		if pop != value {
			return 0
		}

		stack = stack[:len(stack)-1]
	}

	if len(stack) != 0 {
		return 0
	}
	return 1
}