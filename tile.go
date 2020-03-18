package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var n int

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()
	n, _ = strconv.Atoi(text)
	var answer []int64
	answer = make([]int64, n, n)

	for i:=0; i<n; i++{
		if i == 0 {
			answer[0] = 1
			continue
		}
		if i == 1 {
			answer[1] = 2
			continue
		}

		answer[i] = answer[i-1] + answer[i-2]
	}

	fmt.Println(answer[n-1])
}