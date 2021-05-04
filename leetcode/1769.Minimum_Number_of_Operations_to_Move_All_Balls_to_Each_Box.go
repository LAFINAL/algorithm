package leetcode

import "fmt"

func minOperations(boxes string) [] int {
	nonEmpty := make([]int, 0, 0)
	for i:=0; i<len(boxes); i++ {
		if boxes[i] == '1' {
			nonEmpty = append(nonEmpty, i)
		}
	}
	fmt.Println("nonEmpty array: ", nonEmpty)

	answer := make([]int, len(boxes), len(boxes))
	for i:=0; i<len(boxes); i++ {
		operation := 0
		for j:=0; j<len(nonEmpty); j++ {
			distance := nonEmpty[j] - i
			if distance < 0 {
				distance = distance * -1
			}
			operation = operation + distance
		}
		answer[i] = operation
	}
	return answer
}
