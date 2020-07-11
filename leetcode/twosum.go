package main

import "fmt"

func main(){
	a := []int{3, 3}
	answer := twoSum(a, 6)
	answer2 := twoSumWithMap(a, 6)
	fmt.Println(answer)
	fmt.Println(answer2)
}

func twoSum(nums []int, target int) []int {

	answer := make([]int, 2, 2)
	for i:=0; i<len(nums)-1; i++ {
		for j:=i+1; j<len(nums); j++ {
			if nums[i] + nums[j] == target {
				answer[0] = i
				answer[1] = j
			}
		}
	}
	return answer
}

func twoSumWithMap(nums []int, target int) []int {

	answer := make([]int, 2, 2)
	a:=make(map[int]int) // map 선언
	// map으로 만들기 value가 map의 key, index가 map의 value
	for index, value := range(nums) {
		a[value] = index
	}

	complete:=0
	for i:=0; i<len(nums); i++ {
		complete = target-nums[i]
		if a[complete] == 0 || i==a[complete]{
			fmt.Println("test")
			fmt.Println(nums[i], complete)
			continue
		} else {
			answer[0] = i
			answer[1] = a[complete]
			return answer
		}
	}

	return answer
}
