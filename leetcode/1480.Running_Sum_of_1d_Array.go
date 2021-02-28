package leetcode

func runningSum(nums []int) []int {

	for i, _ := range nums {
		if i == 0 {
			continue
		}

		nums[i] = nums[i] + nums[i-1]
	}
	return nums
}
