package main

import (
	"algorithm/leetcode"
	"fmt"
)

func main(){
	a := [][]int{
		{7,0},
		{4,4},
		{7,1},
		{5,0},
		{6,1},
		{5,2},
	}
	answer := leetcode.ReconstructQueue(a)
	fmt.Println("answer: ", answer)
}
