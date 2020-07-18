package leetcode

func SingleNumber(nums []int) int {
	filter := make(map[int]int)
	for i:=0; i<len(nums); i++ {
		value, _ := filter[nums[i]]
		filter[nums[i]] = value + 1
	}

	for key, value := range filter {
		if value == 1 {
			return key
		}
	}

	return 0
}
