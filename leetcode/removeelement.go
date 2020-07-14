package leetcode

import "fmt"

func removeElement(nums []int, val int) int {

	count := 0
	for i:=0; i<len(nums); i++ {
		if nums[i] != val {
			nums[count] = nums[i]
			count += 1
		}
	}
	fmt.Println(nums)

	return count
}
