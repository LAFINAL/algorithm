package leetcode

func MajorityElement(nums []int) int {
	countMap := make(map[int]int)
	for i:=0; i<len(nums); i++ {
		countMap[nums[i]] = countMap[nums[i]] + 1
	}

	for key, value := range countMap {
		if value > len(nums) / 2 {
			return key
		}
	}
	return 0
}