package main

import "fmt"

func main() {
	answer := isValid("{[]}")
	fmt.Println(answer)
}

func isValid(s string) bool {

	stack := make([]byte, 0, 0)
	pare := make(map[byte]byte)
	pare[']'] = '['
	pare['}'] = '{'
	pare[')'] = '('

	for i:=0; i<len(s); i++ {
		if pare[s[i]] == 0 {
			stack = append(stack, s[i])
		} else {
			if len(stack) == 0 {
				return false
			}

			pop := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if pare[s[i]] != pop {
				return false
			}
		}
	}

	if len(stack) != 0 {
		return false
	}

	return true
}