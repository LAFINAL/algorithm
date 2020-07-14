package main

import (
	"algorithm/leetcode"
	"fmt"
)

func main(){
	testCase := []int{3, 2, 4}
	target := 6
	answer := leetcode.TwoSum(testCase, target)
	fmt.Println(answer)
}
