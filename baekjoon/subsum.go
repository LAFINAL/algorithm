package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	buf := make([]byte, 1024*1024)
	scanner.Buffer(buf, 1024*1024)
	scanner.Scan()

	text := scanner.Text()
	splitText := strings.Split(text, " ")
	value, _ := strconv.Atoi(splitText[1])

	scanner.Scan()
	text = scanner.Text()
	elements := strings.Split(text, " ")
	a := make([]int, len(elements), len(elements))
	for i:=0; i<len(elements); i++ {
		a[i], _ = strconv.Atoi(elements[i])
	}

	answer := subsum(a, value)
	fmt.Println(answer)
}

func subsum(a []int, value int) int {

	leftPointer := 0
	rightPointer := 0
	//answerSlice := make([]int, 0, 0)
	answer := len(a)+1
	sum := 0
	for {
		if leftPointer == len(a) {
			break
		}

		if sum >= value {
			if rightPointer - leftPointer < answer {
				answer = rightPointer - leftPointer
			}
			sum = sum - a[leftPointer]
			leftPointer = leftPointer + 1
		} else {
			if rightPointer == len(a) {
				break
			}
			sum = sum + a[rightPointer]
			rightPointer = rightPointer + 1
		}

	}
	if answer == len(a)+1 {
		return 0
	}

	return answer
}