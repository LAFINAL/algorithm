package codility

import "fmt"

func TossReal1() {
	var input string
	_, err := fmt.Scanln(&input)
	if err != nil {
		panic(err)
	}

	if len(input) == 0 {
		fmt.Print("false")
		return
	}

	stack := make([]byte, 0, 0)
	stack = append(stack, input[0])

	for i:=1; i<len(input); i++ {
		top := stack[len(stack)-1]
		if top == '1' && input[i] == '1'{
			fmt.Print("false")
			return
		}
		stack = append(stack, input[i])
	}

	if stack[len(stack)-1] == '1' {
		fmt.Print("false")
		return
	}
	fmt.Print("true")
}
