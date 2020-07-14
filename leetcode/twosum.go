package leetcode

// 짝을 찾는 것 ==> Map을 사용한다.
func TwoSum(nums []int, target int) []int {
	answer := make([]int, 2, 2)

	tempMap := make(map[int]int)
	for index, value := range nums {
		tempMap[value] = index
	}

	complete := 0
	for i:=0; i<len(nums); i++ {
		complete = target - nums[i]
		if value, ok := tempMap[complete]; ok && tempMap[complete] != i {
			answer[0] = i
			answer[1] = value
			break
		}
	}
	return answer
}
