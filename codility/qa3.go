package codility

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TossReal3() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	testCase := strings.Split(input, " ")
	nums := make([]int, len(testCase), len(testCase))
	for i:=0; i<len(testCase); i++ {
		b, _ := strconv.Atoi(testCase[i])
		nums[i] = b
	}

	answer := make([]int, len(nums), len(nums))
	tempMap := make(map[int]int)
	for i:=0; i<len(nums); i++ {
		value, exist := tempMap[nums[i]]
		if !exist {
			value = compute(nums[i])
			tempMap[nums[i]] = value
		}
		answer[i] = value
	}

	for i:=0; i<len(answer); i++ {
		if i != len(answer)-1 {
			fmt.Print(answer[i], " ")
		} else{
			fmt.Print(answer[i])
		}
	}
}

func compute(a int) int {
	if a != 0 {
		fmt.Println("들어옴")
	}
	return 1234
}