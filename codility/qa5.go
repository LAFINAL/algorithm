package codility

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func TossReal5() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	testCase := strings.Split(input, " ")
	kim := make([]int, len(testCase), len(testCase))
	for i:=0; i<len(testCase); i++ {
		b, _ := strconv.Atoi(testCase[i])
		kim[i] = b
	}

	scanner.Scan()
	input = scanner.Text()
	testCase = strings.Split(input, " ")
	lee := make([]int, len(testCase), len(testCase))
	for i:=0; i<len(testCase); i++ {
		b, _ := strconv.Atoi(testCase[i])
		lee[i] = b
	}

	a := make([]int, len(testCase), len(testCase))
	kimSum := 0
	leeSum := 0
	for i:=0; i<len(a); i++ {
		kimSum = kimSum + kim[i]
		leeSum = leeSum + lee[i]
		if kimSum >= leeSum {
			if i!=len(a)-1 {
				fmt.Print(strconv.Itoa(kimSum-leeSum)+" ")
			} else {
				fmt.Print(strconv.Itoa(kimSum-leeSum))
			}
			kimSum = leeSum
		} else {
			if i!=len(a)-1 {
				fmt.Print("0 ")
			} else {
				fmt.Print("0")
			}
		}
	}
}
