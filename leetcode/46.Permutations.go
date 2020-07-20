package leetcode

import "fmt"

func Permute(nums []int) [][]int {

	used := []bool{false, false, false}
	temp := make([]int, 0, 0)
	answer := make([][]int, 0, 0)
	answer = backTrack(nums, used, temp, answer)
	fmt.Println(answer)
	return [][]int{{1,2,3}}
}

func backTrack(nums []int, used []bool, temp []int, answer [][]int) [][]int {
	if len(temp) == len(nums) {
		copySlice := make([]int, len(temp), len(temp))
		copy(copySlice, temp)
		answer = append(answer, copySlice)
		return answer
	}

	for i:=0; i<len(nums); i++ {
		if !used[i] {
			temp = append(temp, nums[i])
			used[i] = true
			answer = backTrack(nums, used, temp, answer)
			used[i] = false
			temp = temp[:len(temp)-1]
		}
	}
	return answer
}
