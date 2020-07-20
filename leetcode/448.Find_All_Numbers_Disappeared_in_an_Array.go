package leetcode

func FindDisappearedNumbers(nums []int) []int {

	check := make([]bool, len(nums), len(nums))
	for i:=0; i<len(nums); i++ {
		check[nums[i]-1] = true
	}

	answer := make([]int, 0, 0)
	for i:=0; i<len(check); i++ {
		if check[i] == false {
			answer = append(answer, i+1)
		}
	}

	return answer
}
