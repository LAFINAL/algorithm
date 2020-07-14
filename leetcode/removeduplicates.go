package leetcode


func removeDuplicates(nums []int) int {

	current := -10
	count := 0
	for i:=0; i<len(nums); i++ {
		if current != nums[i] {
			current = nums[i]
			nums[count] = nums[i]
			count = count + 1
		}
	}
	return count
}
