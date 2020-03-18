package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	n, _ = strconv.Atoi(text)

	var answer []int
	answer = make([]int, n+1, n+1)


	for i:=1; i<n+1; i++ {
		if i == 1 {
			answer[1] = 0
			continue
		}

		if i == 2 {
			answer[2] = 1
			continue
		}
		if i == 3 {
			answer[3] = 1
			continue
		}

		if i%3 == 0 {
			answer[i] = Min(answer[i/3] + 1, answer[i-1] + 1)
			continue
		}

		if i%2 == 0 {
			answer[i] = Min(answer[i/2] + 1, answer[i-1] + 1)
			continue
		}

		answer[i] = answer[i-1] + 1
	}

	fmt.Println(answer[n])
}


func Min(a, b int) (v int){
	if a > b {
		v = b
		return
	}
	v = a
	return
}