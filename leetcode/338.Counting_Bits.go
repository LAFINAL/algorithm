package leetcode

func CountBits(num int) []int {
	answer := make([]int, num+1, num+1)
	for i := 0; i <= num; i++ {
		count := binary(i)
		answer[i] = count
	}
	return answer
}

func binary(num int) int {
	count := 0
	for num/2 != 0 {
		if num%2 == 1{
			count = count + 1
		}
		num = num/2
	}
	if num == 1 {
		count = count + 1
	}
	return count
}
