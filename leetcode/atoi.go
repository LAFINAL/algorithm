package leetcode

import (
	"math"
	"strconv"
	"strings"
)


func myAtoi(str string) int {


	str = strings.TrimSpace(str)

	if len(str)==0{
		return 0
	}

	answerSlice := make([]byte, 0, 0)
	start := 0
	if str[0] == 45 {
		answerSlice = append(answerSlice, str[0])
		start = 1
	}

	if str[0] == 43 {
		start = 1
	}

	for i:=start; i<len(str); i++ {
		if str[i] < 48 || str[i] > 57 {
			break
		}

		answerSlice = append(answerSlice, str[i])
	}

	if len(answerSlice) == 0 || (len(answerSlice) == 1 && answerSlice[0] == 45) {
		return 0
	}
	answer, _ := strconv.Atoi(string(answerSlice))

	if answer > math.MaxInt32 {
		return math.MaxInt32
	}

	if answer < math.MinInt32 {
		return math.MinInt32
	}
	return answer
}